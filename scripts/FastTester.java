import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.PredictionMode;
import org.antlr.v4.runtime.tree.*;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.Path;
import java.util.*;
import java.util.stream.Collectors;
import static java.util.Comparator.reverseOrder;

public class FastTester {
    private static final boolean OVERWRITE_BASELINE = false;
    private static final String BASELINE_FILE = "passed_baseline.txt";

    private static final boolean PRINT_LEAST_HIT_RULES = true;
    private static final int LEAST_HIT_LIMIT = 10; // How many to show

    private static final Map<String, Integer> ruleHitCounts = new HashMap<>();

    public static void main(String[] args) throws Exception {
        performWarmup("/inputs/input.txt");

        System.out.println("-".repeat(95));
        System.out.printf("%-18s | %-12s | %-15s | %-12s | %-15s\n",
                          "File Source", "Success Rate", "Passed / Total", "Total Time", "Avg. Time/Line");
        System.out.println("-".repeat(95));

        testFile("inputs/input.txt", true);
        testFile("inputs/input_full.txt", false);

        System.out.println("-".repeat(95));

        // NEW: Print the dead rule report after all tests run
        reportRuleCoverage();
    }

    private static void testFile(String fileName, boolean checkRegressions) {
        try {
            if (!Files.exists(Paths.get(fileName))) {
                System.out.printf("%-18s | File not found.\n", fileName);
                return;
            }

            List<String> lines = Files.readAllLines(Paths.get(fileName));
            Set<String> currentPasses = new LinkedHashSet<>();
            int total = 0;
            int passed = 0;

            long startTime = System.nanoTime();
            for (String line : lines) {
                String trimmed = line.trim();
                if (trimmed.isEmpty()) continue;
                total++;
                if (runParser(trimmed)) {
                    currentPasses.add(trimmed); // deduplicated, used for regression baseline only
                    passed++;
                }
            }
            long endTime = System.nanoTime();

            if (checkRegressions) {
                handleBaselines(currentPasses);
            }

            long totalDurationNano = endTime - startTime;
            double durationMs = totalDurationNano / 1_000_000.0;
            double avgUsPerLine = (total == 0) ? 0 : (totalDurationNano / 1000.0) / total;
            double rate = (total == 0) ? 0 : (passed / (double) total) * 100;

            String passTotalStr = String.format("%d / %d", passed, total);
            System.out.printf("%-18s | %-8.2f %% | %-15s | %-8.2f ms | %-11.2f us\n",
                fileName, rate, passTotalStr, durationMs, avgUsPerLine);

        } catch (Exception e) {
            System.err.println("Error processing " + fileName + ": " + e.getMessage());
        }
    }

    private static void handleBaselines(Set<String> currentPasses) throws Exception {
        Path path = Paths.get(BASELINE_FILE);

        if (OVERWRITE_BASELINE || !Files.exists(path)) {
            Files.write(path, currentPasses);
            if (OVERWRITE_BASELINE) System.out.println("[INFO] Baseline updated with current passes.");
            return;
        }

        List<String> baseline = Files.readAllLines(path);
        List<String> regressions = baseline.stream()
                .filter(line -> !currentPasses.contains(line))
                .collect(Collectors.toList());

        if (!regressions.isEmpty()) {
            System.err.println("\n" + "!".repeat(20) + " REGRESSION ALERT " + "!".repeat(20));
            System.err.printf("! %d tests that previously passed are now FAILING:\n", regressions.size());
            regressions.forEach(line -> System.err.println("! -> " + line));
            System.err.println("!".repeat(58) + "\n");
        }
    }

    private static void performWarmup(String fileName) {
        try {
            if (!Files.exists(Paths.get(fileName))) return;
            List<String> lines = Files.readAllLines(Paths.get(fileName));
            for (int i = 0; i < Math.min(lines.size(), 500); i++) {
                runParser(lines.get(i));
            }
        } catch (Exception ignored) {}
    }

    private static boolean runParser(String input) {
            RequirementsLexer lexer = new RequirementsLexer(CharStreams.fromString(input));
            lexer.removeErrorListeners();
            RequirementsParser parser = new RequirementsParser(new CommonTokenStream(lexer));
            parser.removeErrorListeners();

            // THE FIX: Custom listener that reports overlaps WITHOUT incrementing syntax errors
            parser.addErrorListener(new BaseErrorListener() {
                @Override
                public void reportAmbiguity(Parser recognizer, org.antlr.v4.runtime.dfa.DFA dfa, int startIndex, int stopIndex, boolean exact, java.util.BitSet ambigAlts, org.antlr.v4.runtime.atn.ATNConfigSet configs) {
                    // You can uncomment the line below to see exactly where the overlaps happen
                    // System.out.println("Rule Overlap (Ambiguity) detected from index " + startIndex + " to " + stopIndex);
                }
            });

            // Note: This mode forces deep lookahead to find ambiguities.
            // If your test suite suddenly feels too slow, comment this line out!
            parser.getInterpreter().setPredictionMode(PredictionMode.LL_EXACT_AMBIG_DETECTION);

            parser.addParseListener(new ParseTreeListener() {
                        @Override
                        public void enterEveryRule(ParserRuleContext ctx) {
                            // Get the rule name and increment its count in the map
                            String ruleName = RequirementsParser.ruleNames[ctx.getRuleIndex()];
                            ruleHitCounts.put(ruleName, ruleHitCounts.getOrDefault(ruleName, 0) + 1);
                        }
                        @Override public void exitEveryRule(ParserRuleContext ctx) {}
                        @Override public void visitTerminal(TerminalNode node) {}
                        @Override public void visitErrorNode(ErrorNode node) {}
                    });

            parser.prog();

            // This will now accurately return true only for actual broken syntax
            return parser.getNumberOfSyntaxErrors() == 0;
        }


    private static void reportRuleCoverage() {
            Set<String> allRules = new HashSet<>(Arrays.asList(RequirementsParser.ruleNames));

            // Dead rules are rules that never made it into our Map
            Set<String> deadRules = new HashSet<>(allRules);
            deadRules.removeAll(ruleHitCounts.keySet());

            System.out.println("\n" + "=".repeat(30) + " RULE COVERAGE REPORT " + "=".repeat(30));
            System.out.printf("Total Grammar Rules: %d\n", allRules.size());
            System.out.printf("Rules Hit: %d\n", ruleHitCounts.size());
            System.out.printf("Dead / Unvisited Rules: %d\n", deadRules.size());

            if (!deadRules.isEmpty()) {
                System.out.println("\nThe following rules were NEVER triggered:");
                deadRules.stream().sorted().forEach(rule -> System.out.println("  - " + rule));
            }

            // NEW: Print the least hit rules
            if (PRINT_LEAST_HIT_RULES && !ruleHitCounts.isEmpty()) {
                System.out.printf("\nLowest Hits (Top/Bottom %d):\n", LEAST_HIT_LIMIT);


                ruleHitCounts.entrySet()
                        .stream().sorted(Map.Entry.comparingByValue(reverseOrder()))
                        .limit(LEAST_HIT_LIMIT / 2 )
                        .forEach(entry -> System.out.printf("  - %-40s : %d hits\n", entry.getKey(), entry.getValue()));

                System.out.printf("   ...\n");
                ruleHitCounts.entrySet().stream()
                        .sorted(Map.Entry.comparingByValue())
                        .limit(LEAST_HIT_LIMIT / 2 )
                        .sorted(Map.Entry.comparingByValue(reverseOrder()))
                        .forEach(entry -> System.out.printf("  - %-40s : %d hits\n", entry.getKey(), entry.getValue()));



            }

            System.out.println("=".repeat(82));
        }
}