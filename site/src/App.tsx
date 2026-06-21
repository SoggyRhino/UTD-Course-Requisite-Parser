import {useMemo, useState} from "react"
import {Card} from "@/components/ui/card"
import UserInfoForm from "@/components/form/UserInfoForm"
import {Input} from "@/components/ui/input"
import requirements from "../../static/requirements.json"
import {Requirements} from "@goscript/parser/objects";
import CourseRequirements from "@/components/evaluation/CourseRequirements.tsx";
import {UserInfo} from "@goscript/parser/objects/constants";
import * as json from "@goscript/encoding/json/index.js";
import * as $ from "@goscript/builtin/index.js";

const coursesList = Object.keys(requirements)
const reqData = requirements as Record<string, any>

function App() {
    const [userInfo, updateUserInfo] = useState<UserInfo>(new UserInfo())
    const [searchedCourse, setSearchedCourse] = useState("")

    const rawEntry = searchedCourse && reqData[searchedCourse] ? reqData[searchedCourse] : null
    const foundData: Requirements | null = useMemo(() => {
        if (!rawEntry) return null
        const req = new Requirements()
        const bytes = $.stringToBytes(JSON.stringify(rawEntry))
        const err = json.Unmarshal(bytes, req)
        if (err != null) {
            console.error("Failed to parse requirements:", err)
            return null
        }
        return req
    }, [JSON.stringify(rawEntry)])

    return (
        <div className="App min-h-screen bg-linear-to-br from-background via-background to-muted/50 p-4 sm:p-8 md:p-12 flex flex-col items-center gap-8">

            <Card className="w-full max-w-3xl border-border/50 shadow-lg bg-card/60 backdrop-blur-xl p-6 sm:p-8 flex flex-col gap-6 rounded-xl transition-all">
                <div className="space-y-1.5">
                    <h2 className="text-xl font-semibold tracking-tight m-0 text-foreground">Course Search</h2>
                    <p className="text-sm text-muted-foreground m-0">Search for a course to evaluate its prerequisites and rules.</p>
                </div>
                <div className="flex flex-col gap-4">
                    <Input
                        placeholder="Search for a course (e.g. CS 1200)"
                        list="app-course-list"
                        value={searchedCourse}
                        onChange={(e) => setSearchedCourse(e.target.value.toUpperCase())}
                        className="h-12 text-base px-4 bg-background/50 focus-visible:ring-primary/50"
                    />
                    <datalist id="app-course-list">
                        {coursesList.map((c) => (
                            <option key={c} value={c}/>
                        ))}
                    </datalist>
                </div>

                <div className="bg-background/40 p-5 rounded-lg border border-border/50 shadow-inner min-h-[120px] overflow-auto">
                    {foundData
                        ? <CourseRequirements info={userInfo} req={foundData}/>
                        : <div className="h-full flex items-center justify-center min-h-[80px]">
                            <p className="text-muted-foreground text-sm font-medium">No course selected</p>
                          </div>
                    }
                </div>
            </Card>

            <Card className="w-full max-w-3xl border-border/50 shadow-lg bg-card/60 backdrop-blur-xl p-6 sm:p-8 flex flex-col gap-6 rounded-xl transition-all">
                <div className="space-y-1.5">
                    <h2 className="text-xl font-semibold tracking-tight m-0 text-foreground">Student Profile</h2>
                    <p className="text-sm text-muted-foreground m-0">Configure your academic information to evaluate requirements.</p>
                </div>
                <div className="pt-2">
                    <UserInfoForm setUserInfo={updateUserInfo} userInfo={userInfo}/>
                </div>
            </Card>
        </div>
    )
}

export default App