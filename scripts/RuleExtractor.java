import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;
import java.nio.file.*;
import java.util.*;
import java.util.stream.Collectors;
import org.antlr.v4.runtime.misc.Interval;

public class RuleExtractor {
    private static final Set<String> TARGET_RULES = new LinkedHashSet<>();

    public static void main(String[] args) throws Exception {

        for (int i =0; i < args.length; i++){
            TARGET_RULES.add(args[i]);
        }

            Files.createDirectories(Paths.get("extracted"));


        String inputFile = "inputs/input.txt";

        if (!Files.exists(Paths.get(inputFile))) {
            System.err.println("input.txt not found.");
            return;
        }

        List<String> lines = Files.readAllLines(Paths.get(inputFile));

        // rule name -> deduplicated ordered set of extracted texts
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

        System.out.printf("Processed %d lines from %s\n", total, inputFile);

        for (Map.Entry<String, LinkedHashSet<String>> entry : results.entrySet()) {
            String rule = entry.getKey();
            LinkedHashSet<String> matches = entry.getValue();
            System.out.printf("\n[%s] — %d unique matches\n", rule, matches.size());
            matches.forEach(m -> System.out.println("  " + m));

            Path out = Paths.get("extracted/extracted_" + rule + ".txt");
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

        // Make sure the token stream is fully loaded so GetTextFromInterval works
        tokens.fill();

        ParseTreeWalker.DEFAULT.walk(new ParseTreeListener() {
            @Override
            public void enterEveryRule(ParserRuleContext ctx) {
                String ruleName = RequirementsParser.ruleNames[ctx.getRuleIndex()];
                if (!results.containsKey(ruleName)) return;

                Token start = ctx.getStart();
                Token stop = ctx.getStop();
                if (start == null || stop == null) return;

                String text = ctx.getStart().getInputStream().getText(
                    new Interval(ctx.getStart().getStartIndex(), ctx.getStop().getStopIndex())
                );
                results.get(ruleName).add(text);
            }

            @Override public void exitEveryRule(ParserRuleContext ctx) {}
            @Override public void visitTerminal(TerminalNode node) {}
            @Override public void visitErrorNode(ErrorNode node) {}
        }, tree);
    }
}
