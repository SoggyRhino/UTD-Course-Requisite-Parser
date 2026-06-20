import { Checkbox } from "@/components/ui/checkbox"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
  FieldLegend,
  FieldSet,
} from "@/components/ui/field"

interface Option {
  label: string
  value: string
}

interface CheckboxGroupFieldProps {
  name: string
  description: string | null
  options: Option[]
  value: string[]
  onChange: (value: string[]) => void
  /**
   * When true, selecting an option clears any other selection — at most one
   * value can be checked. Unlike a radio group, the single checked option
   * can still be unchecked, leaving zero selected (hence "1 or 0").
   */
  singleSelect?: boolean
}

export function CheckboxGroupField({
  name,
  description,
  options,
  value,
  onChange,
  singleSelect = false,
}: CheckboxGroupFieldProps) {
  function toggleOption(option: string, checked: boolean) {
    if (singleSelect) {
      onChange(checked ? [option] : [])
      return
    }
    if (checked) {
      onChange([...value, option])
    } else {
      onChange(value.filter((v) => v !== option))
    }
  }

  return (
    <FieldSet>
      <FieldLegend>{name}</FieldLegend>
      {description && <FieldDescription>{description}</FieldDescription>}
      <FieldGroup data-slot="checkbox-group">
        {options.map((option) => {
          const id = `${name}-${option.value.replace(/\s+/g, "-").toLowerCase()}`
          return (
            <Field key={option.value} orientation="horizontal">
              <Checkbox
                id={id}
                checked={value.includes(option.value)}
                onCheckedChange={(state) =>
                  toggleOption(option.value, state === true)
                }
              />
              <FieldLabel htmlFor={id} className="font-normal">
                {option.label}
              </FieldLabel>
            </Field>
          )
        })}
      </FieldGroup>
    </FieldSet>
  )
}
