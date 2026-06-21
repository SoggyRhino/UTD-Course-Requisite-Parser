import type {Dispatch, SetStateAction} from "react";
import {Field, FieldGroup, FieldLabel} from "@/components/ui/field"
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue,} from "@/components/ui/select"
import {GroupsField} from "./GroupsField"
import {StringInput} from "./StringInput"
import {CourseListField} from "./CourseListField.tsx"

import {
    AnyGrade,
    Freshman,
    Graduate,
    Junior,
    PhD,
    Senior,
    Sophomore,
    Undergraduate,
} from "@goscript/parser/objects/constants/objects.gs"

import type {UserInfo} from "@goscript/parser/objects/constants"


interface UserInfoFormProps {
    setUserInfo:  Dispatch<SetStateAction<UserInfo>>
    userInfo: UserInfo
}


function UserInfoForm({setUserInfo, userInfo}: UserInfoFormProps) {
    const handleGroupsChange = (groups: string[]) => {
        setUserInfo((prev) => {
            const next = prev.clone()
            next.Groups = groups
            return next
        })
    }

    const handleAcademicPlanChange = (plan: string) => {
        setUserInfo((prev) => {
            const next = prev.clone()
            next.AcademicPlan = plan
            return next
        })
    }

    const handleMajorChange = (major: string) => {
        setUserInfo((prev) => {
            const next = prev.clone()
            next.Major = major
            return next
        })
    }

    const isUndergrad = (userInfo.DegreeLevel || Undergraduate) === Undergraduate

    return (
        <FieldGroup className="gap-5">
            <CourseListField userInfo={userInfo} setUserInfo={setUserInfo}/>

            <Field orientation="horizontal">

                <FieldLabel className="w-1/3">Degree
                    level</FieldLabel>
                <Select
                    value={userInfo.DegreeLevel || Undergraduate}
                    onValueChange={(v) =>
                        setUserInfo((prev) => {
                            const next = prev.clone()
                            next.DegreeLevel = v!
                            return next
                        })
                    }
                >
                    <SelectTrigger className="flex-1">
                        <SelectValue placeholder="Select degree level"/>
                    </SelectTrigger>
                    <SelectContent>
                        <SelectItem value={Undergraduate}>Undergraduate</SelectItem>
                        <SelectItem value={Graduate}>Graduate</SelectItem>
                        <SelectItem value={PhD}>PhD</SelectItem>
                    </SelectContent>
                </Select>
            </Field>

            {isUndergrad && (
                <Field orientation="horizontal">
                    <FieldLabel className="w-1/3">Grade
                        level</FieldLabel>
                    <Select
                        value={userInfo.GradeLevel || ""}
                        onValueChange={(v) =>
                            setUserInfo((prev) => {
                                const next = prev.clone()
                                next.GradeLevel = v ?? ""
                                return next
                            })
                        }
                    >
                        <SelectTrigger className="flex-1">
                            <SelectValue placeholder="Select grade level"/>
                        </SelectTrigger>
                        <SelectContent>
                            <SelectItem value={Freshman}>Freshman</SelectItem>
                            <SelectItem value={Sophomore}>Sophomore</SelectItem>
                            <SelectItem value={Junior}>Junior</SelectItem>
                            <SelectItem value={Senior}>Senior</SelectItem>
                            <SelectItem value={AnyGrade}>Any level</SelectItem>
                        </SelectContent>
                    </Select>
                </Field>
            )}

            <StringInput
                name="Academic Plan"
                value={userInfo.AcademicPlan ?? ""}
                onChange={handleAcademicPlanChange}
            />

            <StringInput
                name="Major"
                value={userInfo.Major ?? ""}
                onChange={handleMajorChange}
            />

            <GroupsField value={userInfo.Groups ?? []} onChange={handleGroupsChange}/>
        </FieldGroup>
    )
}

export default UserInfoForm