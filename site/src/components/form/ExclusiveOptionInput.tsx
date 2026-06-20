import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
  FieldLegend,
  FieldSet,
} from "@/components/ui/field"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"

interface Option {
  label: string
  value: string
}

interface ExclusiveOptionInputProps {
  name: string
  description: string | null
  value: string
  defaultValue: string
  options: Option[]
  onChange: (value: string) => void
}

export function ExclusiveOptionInput({
  name,
  description,
  value,
  defaultValue,
  options,
  onChange,
}: ExclusiveOptionInputProps) {
  return (
    <FieldSet>
      <FieldLegend>{name}</FieldLegend>
      {description && <FieldDescription>{description}</FieldDescription>}
      <FieldGroup>
        <RadioGroup
          value={value || defaultValue}
          onValueChange={onChange}
        >
          {options.map((option) => {
            const id = `${name}-${option.value.replace(/\s+/g, "-").toLowerCase()}`
            return (
              <Field key={option.value} orientation="horizontal">
                <RadioGroupItem id={id} value={option.value} />
                <FieldLabel htmlFor={id} className="font-normal">
                  {option.label}
                </FieldLabel>
              </Field>
            )
          })}
        </RadioGroup>
      </FieldGroup>
    </FieldSet>
  )
}
