import { Field, FieldLabel } from "@/components/ui/field.tsx";
import {Input} from "@/components/ui/input.tsx";

interface StringInputProps {
    name: string
    value: string
    onChange: (str: string) => void
}

export function StringInput({name, value, onChange}: StringInputProps) {
    return (
        <Field orientation="horizontal">
            <FieldLabel className="w-1/3">{name}</FieldLabel>
            <Input placeholder="" value={value} onChange={(e) => onChange(e.target.value)} className="flex-1" />
        </Field>
    )
}
