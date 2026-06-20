import { Field, FieldLabel } from "@/components/ui/field"
import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuGroup,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Plus, X } from "lucide-react"

// Mirrors the StudentGroup constants from constants.go.
// NOTE: only the ones explicitly typed as StudentGroup in the Go source
// (ComputerScholarsProgram) have a real string constant on that type today —
// CollegeVHonors / LiberalArtsHonors / SCVG / DMHP / DLAH are declared as
// plain `string`, not `StudentGroup`. Listing their literal values here so
// the picker covers all known groups; if Go promotes them to StudentGroup
// later, swap these literals for the exported consts.
const KNOWN_STUDENT_GROUPS: { label: string; value: string }[] = [
  { label: "Computer Scholars Program", value: "Computer Scholars Program" },
  { label: "Collegium V Honors", value: "Collegium V Honors" },
  { label: "Liberal Arts Honors", value: "Liberal Arts Honors" },
  { label: "SCVG", value: "SCVG" },
  { label: "DMHP", value: "DMHP" },
  { label: "DLAH", value: "DLAH" },
]

interface GroupsFieldProps {
  value: string[]
  onChange: (groups: string[]) => void
}

export function GroupsField({ value, onChange }: GroupsFieldProps) {
  function toggleGroup(group: string, checked: boolean) {
    if (checked) {
      onChange([...value, group])
    } else {
      onChange(value.filter((g) => g !== group))
    }
  }

  return (
    <Field orientation="horizontal">
      <FieldLabel className="w-1/3 sm:w-auto min-w-[100px] self-start mt-2">Groups</FieldLabel>
      <div className="flex flex-1 flex-wrap items-center gap-2">
        {value.map((groupVal) => {
          const groupInfo = KNOWN_STUDENT_GROUPS.find(g => g.value === groupVal)
          const label = groupInfo ? groupInfo.label : groupVal
          return (
            <div
              key={groupVal}
              className="flex items-center gap-1 rounded-md border bg-secondary/50 px-2.5 py-1 text-xs font-medium text-secondary-foreground transition-colors"
            >
              <span>{label}</span>
              <button
                type="button"
                className="ml-1 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2"
                onClick={() => toggleGroup(groupVal, false)}
                aria-label={`Remove ${label}`}
              >
                <X className="size-3" />
              </button>
            </div>
          )
        })}

        <DropdownMenu>
          <DropdownMenuTrigger render={<Button variant="outline" size="sm" className="h-7 border-dashed gap-1 rounded-md text-xs" />}>
            <Plus className="size-3.5" />
            Add group
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-56" align="start">
            <DropdownMenuGroup>
              <DropdownMenuLabel>Student Groups</DropdownMenuLabel>
              <DropdownMenuSeparator />
              {KNOWN_STUDENT_GROUPS.map((group) => (
                <DropdownMenuCheckboxItem
                  key={group.value}
                  checked={value.includes(group.value)}
                  onCheckedChange={(state) => toggleGroup(group.value, state)}
                >
                  {group.label}
                </DropdownMenuCheckboxItem>
              ))}
            </DropdownMenuGroup>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </Field>
  )
}
