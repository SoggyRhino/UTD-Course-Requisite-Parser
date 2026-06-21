import {Requirements, RequirementsResult} from "@goscript/parser/objects"
import {type Evaluation, UserInfo} from "@goscript/parser/objects/constants"
import {EvaluationNode} from "./EvaluationNode.tsx"
import {cn} from "@/lib/utils.ts"

type CourseRequirementsProps = {
    info: UserInfo
    req: Requirements
}

const OVERALL_STYLES: Record<string, { text: string; label: string }> = {
    pass: {text: "text-emerald-700 dark:text-emerald-400", label: "Requirements met"},
    definite_fail: {text: "text-red-700 dark:text-red-400", label: "Requirements not met"},
    possible_fail: {text: "text-amber-700 dark:text-amber-400", label: "May not meet requirements"},
    unknown: {text: "text-muted-foreground", label: "Status unknown"},
    invalid_rule: {text: "text-purple-700 dark:text-purple-400", label: "Invalid rule"},
    system_error: {text: "text-red-800 dark:text-red-500", label: "System error"},
}

function unwrapEvaluation(value: unknown): Evaluation | null {
    if (value == null) return null
    if (typeof value === "object" && "value" in (value as Record<string, unknown>)) {
        const inner = (value as { value: unknown }).value
        return (inner ?? null) as Evaluation | null
    }
    return value as Evaluation
}

interface SectionProps {
    title: string
    evaluations: Evaluation[]
}

function Section({title, evaluations}: SectionProps) {
    if (evaluations.length === 0) return null
    return (
        <div>
            <h3 className="text-xs font-medium uppercase tracking-wide text-muted-foreground mb-1.5 px-2">
                {title}
            </h3>
            <div className="rounded-md border bg-card divide-y divide-border/60">
                {evaluations.map((evaluation, i) => (
                    <EvaluationNode key={`${evaluation.Name}-${i}`} evaluation={evaluation}/>
                ))}
            </div>
        </div>
    )
}

function CourseRequirements({info, req}: CourseRequirementsProps) {
    const result: RequirementsResult = req.Evaluate(info)

    const preReqs = unwrapEvaluation(result.PreReqs)
    const coReqs = unwrapEvaluation(result.CoReqs)
    const preOrCoReqs = unwrapEvaluation(result.PreOrCoReqs)
    const rules = result.Rules ?? []
    const notices = result.Notices ?? []

    const preReqsList = preReqs ? [preReqs] : []
    const coReqsList = coReqs ? [coReqs] : []
    const preOrCoReqsList = preOrCoReqs ? [preOrCoReqs] : []

    const overall = OVERALL_STYLES[result.Overall] ?? {
        text: "text-muted-foreground",
        label: result.Overall || "Unknown",
    }

    const hasAnything =
        preReqsList.length + coReqsList.length + preOrCoReqsList.length + rules.length + notices.length > 0

    return (
        <div className="space-y-4">
            <div className="flex items-center justify-between rounded-md border bg-muted/30 px-3 py-2">
                <span className="text-sm font-medium">Overall</span>
                <span className={cn("text-sm font-medium", overall.text)}>{overall.label}</span>
            </div>

            {!hasAnything && (
                <p className="text-sm text-muted-foreground px-2">
                    This course has no requirements to evaluate.
                </p>
            )}

            <Section title="Prerequisites" evaluations={preReqsList}/>
            <Section title="Corequisites" evaluations={coReqsList}/>
            <Section title="Prerequisites or corequisites" evaluations={preOrCoReqsList}/>
            <Section title="Rules" evaluations={rules}/>
            <Section title="Notices" evaluations={notices}/>
        </div>
    )
}

export default CourseRequirements