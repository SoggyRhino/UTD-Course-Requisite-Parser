import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;
import java.nio.file.*;
import java.util.*;
import java.util.stream.Collectors;
import org.antlr.v4.runtime.misc.Interval;

public class RuleExtractor {
    private static final Set<String> TARGET_RULES = new LinkedHashSet<>();
    private static final String INPUT = "inputs/input.txt";
    public static void main(String[] args) throws Exception {

        List<String> rules = new ArrayList<>();

        for (String arg : args) {
            TARGET_RULES.add(arg);
        }

        Files.createDirectories(Paths.get("out"));

        List<String> lines = Files.readAllLines(Paths.get(INPUT));
        Map<String, LinkedHashSet<String>> results = new LinkedHashMap<>();
        for (String rule : TARGET_RULES) {
            results.put(rule, new LinkedHashSet<>());
        }

        int total = 0;
        for (String line : lines) {
            String trimmed = line.trim();
            if (trimmed.isEmpty()) continue;
            total++;
            extractFromLine(trimmed, results);
        }

        for (Map.Entry<String, LinkedHashSet<String>> entry : results.entrySet()) {
            String rule = entry.getKey();
            LinkedHashSet<String> matches = entry.getValue();
            System.out.printf("\n[%s] — %d unique matches\n", rule, matches.size());
            matches.forEach(m -> System.out.println("  " + m));

            Path out = Paths.get("out/extracted_" + rule + ".txt");
            Files.write(out, matches);
            System.out.printf("  -> Written to %s", out);
        }

    }

    private static void extractFromLine(String input, Map<String, LinkedHashSet<String>> results) {
        RequirementsLexer lexer = new RequirementsLexer(CharStreams.fromString(input));
        lexer.removeErrorListeners();
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        RequirementsParser parser = new RequirementsParser(tokens);
        parser.removeErrorListeners();

        ParseTree tree = parser.prog();

        tokens.fill();

        ParseTreeWalker.DEFAULT.walk(new ParseTreeListener() {
            @Override
            public void enterEveryRule(ParserRuleContext ctx) {
                String ruleName = RequirementsParser.ruleNames[ctx.getRuleIndex()];

                String className = ctx.getClass().getSimpleName();
                String labelName = className.endsWith("Context")
                    ? className.substring(0, className.length() - 7)
                    : className;

                String formattedLabel = Character.toLowerCase(labelName.charAt(0)) + labelName.substring(1);

                String matchFound = null;

                if (results.containsKey(formattedLabel)) {
                    matchFound = formattedLabel;
                } else if (results.containsKey(ruleName)) {
                    matchFound = ruleName;
                }

                if (matchFound == null) return;

                int a = ctx.getStart().getStartIndex();
                int b = ctx.getStop().getStopIndex();
                Interval interval = new Interval(a, b);
                String text = ctx.getStart().getInputStream().getText(interval);

                results.get(matchFound).add(text);
            }

            @Override public void exitEveryRule(ParserRuleContext ctx) {}
            @Override public void visitTerminal(TerminalNode node) {}
            @Override public void visitErrorNode(ErrorNode node) {}
        }, tree);
    }
}
