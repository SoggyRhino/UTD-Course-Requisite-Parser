import { Trash2 } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Checkbox } from "@/components/ui/checkbox"
import {
  Tooltip,
  TooltipContent, TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import type { Grade } from "@goscript/parser/objects/constants/objects.gs"
import requirements from "../../../../static/requirements.json"

const GRADE_OPTIONS: Grade[] = [
  "A+", "A", "A-",
  "B+", "B", "B-",
  "C+", "C", "C-",
  "D+", "D", "D-",
  "F",
]

const coursesMap = new Map<string, { prefix: string; number: string }>()
Object.keys(requirements).forEach((course) => {
  const [prefix, number] = course.split(" ")
  coursesMap.set(course.replace(/\s+/g, ""), { prefix, number })
})

const coursesList = Object.keys(requirements)

export interface CourseRowState {
  id: string
  rawInput: string
  prefix: string
  number: string
  grade: Grade | ""
  currentlyEnrolled: boolean
}

interface CourseRowProps {
  course: CourseRowState
  onChange: (course: CourseRowState) => void
  onRemove: () => void
}

export function CourseRow({
  course,
  onChange,
  onRemove,
}: CourseRowProps) {
  function update(patch: Partial<CourseRowState>) {
    onChange({ ...course, ...patch })
  }

  function handleEnrolledChange(checked: boolean) {
    update({ currentlyEnrolled: checked, grade: checked ? "" : course.grade })
  }

  const isValid = course.prefix && course.number

  return (
    <div className="flex items-center gap-4 w-full">
      <div className="flex-1">
        <Input
          placeholder="CS 1234"
          list="course-list"
          value={course.rawInput}
          onChange={(e) => {
            const val = e.target.value
            const searchStr = val.toUpperCase().replace(/\s+/g, "")
            const matched = coursesMap.get(searchStr)
            
            const regex = /^[a-zA-Z]{2,4}\s[0-9Vv]{4}$/
            const passesRegex = regex.test(val)

            if (matched) {
              update({ rawInput: val, prefix: matched.prefix, number: matched.number })
            } else if (passesRegex) {
              const parts = val.split(/\s+/)
              update({ rawInput: val, prefix: parts[0].toUpperCase(), number: parts[1].toUpperCase() })
            } else {
              update({ rawInput: val, prefix: "", number: "" })
            }
          }}
          className={!isValid && course.rawInput.length > 0 ? "border-red-500 focus-visible:ring-red-500" : ""}
        />
        <datalist id="course-list">
          {coursesList.map((c) => (
            <option key={c} value={c} />
          ))}
        </datalist>
      </div>

      <div className="w-[120px]">
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger render={<span onPointerDown={(e) => e.preventDefault()} />} closeOnClick={false}>
                <Select
                  value={course.grade}
                  onValueChange={(v) => update({ grade: v as Grade })}
                  disabled={course.currentlyEnrolled}
                >
                  <SelectTrigger className="w-full" aria-label="Grade">
                    <SelectValue placeholder="Grade" />
                  </SelectTrigger>
                  <SelectContent>
                    {GRADE_OPTIONS.map((g) => (
                      <SelectItem key={g} value={g}>
                        {g}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
            </TooltipTrigger>
            {course.currentlyEnrolled && (
              <TooltipContent>
                You can&apos;t set a grade for a course you&apos;re currently enrolled in.
              </TooltipContent>
            )}
          </Tooltip>
        </TooltipProvider>
      </div>

      <div className="flex items-center gap-2 w-[110px]">
        <Checkbox
          id={`${course.id}-enrolled`}
          checked={course.currentlyEnrolled}
          onCheckedChange={(state) => handleEnrolledChange(state === true)}
        />
        <label
          htmlFor={`${course.id}-enrolled`}
          className="text-sm font-normal cursor-pointer select-none"
        >
          Enrolled
        </label>
      </div>

      <Button
        type="button"
        variant="ghost"
        size="icon"
        aria-label="Remove course"
        onClick={onRemove}
      >
        <Trash2 className="size-4" />
      </Button>
    </div>
  )
}
