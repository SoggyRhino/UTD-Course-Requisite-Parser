import { useState } from "react"
import { ChevronRight } from "lucide-react"
import {
    type EvalStatus,
    type Evaluation,
    StatusPass,
    StatusDefiniteFail,
    StatusPossibleFail,
    StatusUnknown,
    StatusInvalidRule,
    StatusSystemError,
} from "@goscript/parser/objects/constants"
import { cn } from "@/lib/utils.ts"

// Visual language for each EvalStatus. Kept to shadcn-style semantic colors
// (text/bg/border triplets) rather than a custom palette, so this slots into
// the existing muted UI instead of introducing a new accent system.
const STATUS_STYLES: Record<
    EvalStatus,
    { dot: string; text: string; label: string }
> = {
    [StatusPass]: {
        dot: "bg-emerald-500",
        text: "text-emerald-700 dark:text-emerald-400",
        label: "Pass",
    },
    [StatusDefiniteFail]: {
        dot: "bg-red-500",
        text: "text-red-700 dark:text-red-400",
        label: "Fail",
    },
    [StatusPossibleFail]: {
        dot: "bg-amber-500",
        text: "text-amber-700 dark:text-amber-400",
        label: "Possible fail",
    },
    [StatusUnknown]: {
        dot: "bg-muted-foreground/50",
        text: "text-muted-foreground",
        label: "Unknown",
    },
    [StatusInvalidRule]: {
        dot: "bg-purple-500",
        text: "text-purple-700 dark:text-purple-400",
        label: "Invalid rule",
    },
    [StatusSystemError]: {
        dot: "bg-red-700",
        text: "text-red-800 dark:text-red-500",
        label: "System error",
    },
}

function statusStyle(status: EvalStatus) {
    return (
        STATUS_STYLES[status] ?? {
            dot: "bg-muted-foreground/50",
            text: "text-muted-foreground",
            label: status || "Unknown",
        }
    )
}

interface EvaluationNodeProps {
    evaluation: Evaluation
    depth?: number
    defaultOpen?: boolean
}

export function EvaluationNode({ evaluation, depth = 0, defaultOpen }: EvaluationNodeProps) {
    const children = evaluation.Children ?? []
    const hasChildren = children.length > 0
    const style = statusStyle(evaluation.Status)

    // Auto-expand failing/unresolved branches so the user lands on the
    // problem without extra clicks; passing branches stay collapsed.
    const [isOpen, setIsOpen] = useState(
        defaultOpen ?? evaluation.Status !== StatusPass
    )

    return (
        <div className="text-sm">
            <button
                type="button"
                onClick={() => hasChildren && setIsOpen((o) => !o)}
                className={cn(
                    "flex w-full items-start gap-2 rounded-md px-2 py-1.5 text-left transition-colors",
                    hasChildren && "hover:bg-muted/60 cursor-pointer",
                    !hasChildren && "cursor-default"
                )}
                style={{ paddingLeft: `${depth * 20 + 8}px` }}
            >
                {hasChildren ? (
                    <ChevronRight
                        className={cn(
                            "size-3.5 mt-0.5 shrink-0 text-muted-foreground transition-transform",
                            isOpen && "rotate-90"
                        )}
                    />
                ) : (
                    <span className="size-3.5 shrink-0" />
                )}

                <span className={cn("mt-1.5 size-1.5 shrink-0 rounded-full", style.dot)} />

                <span className="flex-1 min-w-0">
                    <span className="font-medium text-foreground">{evaluation.Name}</span>
                    {evaluation.Summary && (
                        <span className="block text-muted-foreground text-xs mt-0.5">
                            {evaluation.Summary}
                        </span>
                    )}
                </span>

                <span className={cn("shrink-0 text-xs font-medium", style.text)}>
                    {style.label}
                </span>
            </button>

            {hasChildren && isOpen && (
                <div>
                    {children.map((child, i) => (
                        <EvaluationNode key={`${child.Name}-${i}`} evaluation={child} depth={depth + 1} />
                    ))}
                </div>
            )}
        </div>
    )
}
