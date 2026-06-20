import { useState } from "react"
import { Card, CardContent, CardTitle } from "@/components/ui/card"
import {
    Collapsible,
    CollapsibleContent,
    CollapsibleTrigger,
} from "@/components/ui/collapsible"
import { ChevronDown } from "lucide-react"
import { UserInfo } from "@goscript/parser/objects/constants/objects.gs"
import UserInfoForm from "@/components/form/UserInfoForm"
import { Input } from "@/components/ui/input"
import requirements from "../../static/requirements.json"
import {Requirements} from "@goscript/parser/objects";
import CourseRequirements from "@/components/CourseRequirements.tsx";
const coursesList = Object.keys(requirements)
const reqData = requirements as Record<string, any>

function App() {
    const [userInfo, updateUserInfo] = useState<UserInfo>(new UserInfo())
    const [isOpen, setIsOpen] = useState(true)
    const [searchedCourse, setSearchedCourse] = useState("")

    const foundData: Requirements | null  = searchedCourse && reqData[searchedCourse] ? reqData[searchedCourse] : null

    return (
        <div className="App min-h-screen bg-muted/30 p-4 sm:p-8 md:p-12 flex flex-col items-center gap-6">
            <Collapsible open={isOpen} onOpenChange={setIsOpen} className="w-full max-w-3xl">
                <Card className="border-muted/50 shadow-sm">
                    <CollapsibleTrigger className="flex w-full items-center justify-between border-b px-6 py-5 text-left">
                        <CardTitle className="text-lg font-medium tracking-tight">
                            Student information
                        </CardTitle>
                        <ChevronDown
                            className={`size-4 shrink-0 text-muted-foreground transition-transform ${isOpen ? "rotate-180" : ""}`}
                        />
                    </CollapsibleTrigger>
                    <CollapsibleContent>
                        <CardContent >
                            <UserInfoForm setUserInfo={updateUserInfo} userInfo={userInfo} />
                        </CardContent>
                    </CollapsibleContent>
                </Card>
            </Collapsible>

            <Card className="w-full max-w-3xl border-muted/50 shadow-sm p-6 flex flex-col gap-4">
                <h2 className="text-lg font-medium tracking-tight m-0">Course Search</h2>
                <Input 
                    placeholder="Search for a course (e.g. CS 1200)"
                    list="app-course-list"
                    value={searchedCourse}
                    onChange={(e) => setSearchedCourse(e.target.value.toUpperCase())}
                />
                <datalist id="app-course-list">
                    {coursesList.map((c) => (
                        <option key={c} value={c} />
                    ))}
                </datalist>

                <div className="bg-muted/30 p-4 rounded-md border min-h-[100px] overflow-auto">
                    {foundData ? (
                        <CourseRequirements req={foundData}/>
                    ) : (
                        <p className="text-sm text-muted-foreground">
                            {searchedCourse ? "Course not found in requirements." : "Enter a course to view its requirements JSON."}
                        </p>
                    )}
                </div>
            </Card>
        </div>
    )
}

export default App