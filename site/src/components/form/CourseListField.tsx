import { useEffect, useMemo, useState } from "react"
import { ChevronDown, Plus } from "lucide-react"
import { Button } from "@/components/ui/button"
import { FieldGroup, FieldLegend, FieldSet } from "@/components/ui/field"
import { Course, type Grade, type UserInfo } from "@goscript/parser/objects/constants"
import { CourseRow, type CourseRowState } from "./CourseRow"
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible"

export interface CourseListFieldProps {
  userInfo: UserInfo
  setUserInfo: (value: ((prevState: UserInfo) => UserInfo) | UserInfo) => void
}

let nextId = 0
function makeId() {
  nextId += 1
  return `course-${nextId}`
}

function deriveRows(userInfo: UserInfo): CourseRowState[] {
  const rows: CourseRowState[] = []

  const taken = userInfo.Taken
  if (taken) {
    for (const [course, grade] of taken) {
      rows.push({
        id: `taken:${course.Prefix}-${course.Number}-${course.Section}`,
        rawInput: `${course.Prefix} ${course.Number}`,
        prefix: course.Prefix,
        number: course.Number,
        grade: grade as Grade,
        currentlyEnrolled: false,
      })
    }
  }

  const enrolled = userInfo.CurrentEnrollment ?? []
  for (const course of enrolled) {
    rows.push({
      id: `enrolled:${course.Prefix}-${course.Number}-${course.Section}`,
      rawInput: `${course.Prefix} ${course.Number}`,
      prefix: course.Prefix,
      number: course.Number,
      grade: "",
      currentlyEnrolled: true,
    })
  }

  return rows
}

export function CourseListField({ userInfo, setUserInfo }: CourseListFieldProps) {
  const [draftRows, setDraftRows] = useState<CourseRowState[]>([])
  const [isOpen, setIsOpen] = useState(false)

  const committedRows = useMemo(() => {
    return deriveRows(userInfo).filter((cr) => 
      !draftRows.some((dr) => 
        dr.prefix === cr.prefix && dr.number === cr.number && dr.prefix !== "" && dr.number !== ""
      )
    )
  }, [userInfo, draftRows])

  const rows = [...committedRows, ...draftRows]

  useEffect(() => {
    if (rows.length === 0) {
      setDraftRows([{ id: makeId(), rawInput: "", prefix: "", number: "", grade: "", currentlyEnrolled: false }])
      setIsOpen(true)
    }
  }, [rows.length])

  function isDraft(id: string) {
    return draftRows.some((r) => r.id === id)
  }

  function writeBack(oldRow: CourseRowState | undefined, course: CourseRowState) {
    setUserInfo((prev) => {
      const next = prev.clone()

      const key = new Course({
        Prefix: course.prefix,
        Number: course.number,
        Section: "",
      })

      const nextTaken = new Map(next.Taken ?? [])

      if (oldRow && oldRow.prefix && oldRow.number && (oldRow.prefix !== course.prefix || oldRow.number !== course.number)) {
        for (const k of nextTaken.keys()) {
          if (k.Prefix === oldRow.prefix && k.Number === oldRow.number) {
            nextTaken.delete(k)
          }
        }
        next.CurrentEnrollment = (next.CurrentEnrollment ?? []).filter(
          (c) => !(c.Prefix === oldRow.prefix && c.Number === oldRow.number)
        )
      }

      if (course.prefix && course.number) {
        for (const k of nextTaken.keys()) {
          if (k.Prefix === course.prefix && k.Number === course.number) {
            nextTaken.delete(k)
          }
        }
        next.CurrentEnrollment = (next.CurrentEnrollment ?? []).filter(
          (c) => !(c.Prefix === course.prefix && c.Number === course.number)
        )
      }

      if (course.prefix && course.number) {
        if (course.currentlyEnrolled) {
          next.CurrentEnrollment = [...(next.CurrentEnrollment ?? []), key]
        } else if (course.grade) {
          nextTaken.set(key, course.grade)
        }
      }

      next.Taken = nextTaken
      return next
    })
  }

  function handleRowChange(id: string, updated: CourseRowState) {
    const isDraftRow = isDraft(id)
    const oldRow = isDraftRow ? draftRows.find((r) => r.id === id) : committedRows.find((r) => r.id === id)

    if (isDraftRow) {
      setDraftRows((prev) => prev.map((r) => (r.id === id ? updated : r)))
    } else {
      setDraftRows((prev) => [...prev, updated])
    }

    writeBack(oldRow, updated)
  }

  function handleRowRemove(id: string, course: CourseRowState) {
    setDraftRows((prev) => prev.filter((r) => r.id !== id))
    
    setUserInfo((prev) => {
      const next = prev.clone()
      const nextTaken = new Map(next.Taken ?? [])
      for (const k of nextTaken.keys()) {
        if (k.Prefix === course.prefix && k.Number === course.number) {
          nextTaken.delete(k)
        }
      }
      next.Taken = nextTaken
      next.CurrentEnrollment = (next.CurrentEnrollment ?? []).filter(
        (c) => !(c.Prefix === course.prefix && c.Number === course.number)
      )
      return next
    })
  }

  function handleAddCourse() {
    setDraftRows((prev) => [
      ...prev,
      { id: makeId(), rawInput: "", prefix: "", number: "", grade: "", currentlyEnrolled: false },
    ])
    setIsOpen(true)
  }

  return (
    <FieldSet className="w-full">
      <Collapsible open={isOpen} onOpenChange={setIsOpen} className="rounded-md border bg-card">
        <CollapsibleTrigger className="flex w-full cursor-pointer items-center justify-between p-4 hover:bg-muted/50 transition-colors">
            <div className="text-left">
              <FieldLegend className="mb-0">Courses</FieldLegend>
            </div>
            <ChevronDown className={`size-5 shrink-0 text-muted-foreground transition-transform ${isOpen ? "rotate-180" : ""}`} />
        </CollapsibleTrigger>
        
        <CollapsibleContent className="px-4 pb-4">
          <FieldGroup className="gap-3">
            {rows.map((row) => (
              <CourseRow
                key={row.id}
                course={row}
                onChange={(updated) => handleRowChange(row.id, updated)}
                onRemove={() => handleRowRemove(row.id, row)}
              />
            ))}

            <Button type="button" variant="outline" onClick={handleAddCourse} className="w-fit">
              <Plus className="size-4" />
              Add course
            </Button>
          </FieldGroup>
        </CollapsibleContent>
      </Collapsible>
    </FieldSet>
  )
}
