import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.PredictionMode;
import org.antlr.v4.runtime.tree.*;

import java.nio.file.*;
import java.util.*;
import java.util.stream.Collectors;

import static java.util.Comparator.reverseOrder;

public class FastTester {
    private static String baselineFile = "out/passed_baseline.txt";
    private static final Map<String, Integer> ruleHitCounts = new HashMap<>();

    public static void main(String[] args) throws Exception {
        boolean overwriteBaseline = Arrays.asList(args).contains("--overwrite-baseline");

        int leastHitsLimit = 10;
        String warmupFile = "inputs/input.txt";
        List<String> inputFiles = new ArrayList<>();

        for (String arg : args) {
            if (arg.startsWith("--least-hits=")) {
                leastHitsLimit = Integer.parseInt(arg.substring("--least-hits=".length()));
            } else if (!arg.startsWith("--")) {
                inputFiles.add(arg);
            }
        }

        Files.createDirectories(Paths.get("out"));
        performWarmup(warmupFile);

        System.out.println("-".repeat(95));
        System.out.printf("%-25s | %-12s | %-15s | %-12s | %-15s\n",
                "File Source", "Success Rate", "Passed / Total", "Total Time", "Avg. Time/Line");
        System.out.println("-".repeat(95));


        if (inputFiles.isEmpty()) {
            inputFiles.add("inputs/input.txt");
            inputFiles.add("inputs/input_full.txt");
        }

        for (int i = 0; i < inputFiles.size(); i++) {
            testFile(inputFiles.get(i), i == 0, overwriteBaseline);
        }

        System.out.println("-".repeat(95));
        reportRuleCoverage(leastHitsLimit);

    }

    private static void testFile(String fileName, boolean checkRegressions, boolean overwriteBaseline) {
        try {
            if (!Files.exists(Paths.get(fileName))) {
                System.out.printf("%-25s | File not found.\n", fileName);
                return;
            }
            List<String> lines = Files.readAllLines(Paths.get(fileName));
            Set<String> currentPasses = new LinkedHashSet<>();
            int total = 0, passed = 0;

            long startTime = System.nanoTime();
            for (String line : lines) {
                String trimmed = line.trim();
                if (trimmed.isEmpty()) continue;
                total++;
                if (runParser(trimmed)) {
                    currentPasses.add(trimmed);
                    passed++;
                }
            }
            long endTime = System.nanoTime();

            if (checkRegressions) {
                handleBaselines(currentPasses, overwriteBaseline);
            }

            long totalNano = endTime - startTime;
            double durationMs = totalNano / 1_000_000.0;
            double avgUs = total == 0 ? 0 : (totalNano / 1000.0) / total;
            double rate = total == 0 ? 0 : (passed / (double) total) * 100;

            System.out.printf("%-25s | %-8.2f %% | %-15s | %-8.2f ms | %-11.2f us\n", fileName, rate, passed + " / " + total, durationMs, avgUs);
        } catch (Exception e) {
            System.err.println("Error processing " + fileName + ": " + e.getMessage());
        }
    }

    private static void handleBaselines(Set<String> currentPasses, boolean overwriteBaseline) throws Exception {
        Path path = Paths.get(baselineFile);

        if (overwriteBaseline || !Files.exists(path)) {
            Files.write(path, currentPasses);
            if (overwriteBaseline) System.out.println("[INFO] Baseline updated with current passes.");
            return;
        }

        List<String> regressions = Files.readAllLines(path).stream().filter(line -> !currentPasses.contains(line)).collect(Collectors.toList());

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
        } catch (Exception ignored) {
        }
    }

    private static boolean runParser(String input) {
        RequirementsLexer lexer = new RequirementsLexer(CharStreams.fromString(input));
        lexer.removeErrorListeners();

        RequirementsParser parser = new RequirementsParser(new CommonTokenStream(lexer));
        parser.removeErrorListeners();

        parser.addErrorListener(new BaseErrorListener() {
            @Override
            public void reportAmbiguity(Parser recognizer, org.antlr.v4.runtime.dfa.DFA dfa, int startIndex, int stopIndex, boolean exact, java.util.BitSet ambigAlts, org.antlr.v4.runtime.atn.ATNConfigSet configs) {

            }
        });

        parser.getInterpreter().setPredictionMode(PredictionMode.LL_EXACT_AMBIG_DETECTION);

        parser.addParseListener(new ParseTreeListener() {
            @Override
            public void enterEveryRule(ParserRuleContext ctx) {
                String rule = RequirementsParser.ruleNames[ctx.getRuleIndex()];
                ruleHitCounts.merge(rule, 1, Integer::sum);
            }

            @Override
            public void exitEveryRule(ParserRuleContext ctx) {
            }

            @Override
            public void visitTerminal(TerminalNode node) {
            }

            @Override
            public void visitErrorNode(ErrorNode node) {
            }
        });

        parser.prog();
        return parser.getNumberOfSyntaxErrors() == 0;
    }

    private static void reportRuleCoverage(int leastHitsLimit) {
        Set<String> allRules = new HashSet<>(Arrays.asList(RequirementsParser.ruleNames));

        Set<String> deadRules = new HashSet<>(allRules);
        deadRules.removeAll(ruleHitCounts.keySet());

        System.out.println("\n" + "=".repeat(30) + " RULE COVERAGE REPORT " + "=".repeat(30));
        System.out.printf("Total Grammar Rules:    %d\n", allRules.size());
        System.out.printf("Rules Hit:              %d\n", ruleHitCounts.size());
        System.out.printf("Dead / Unvisited Rules: %d\n", deadRules.size());

        if (!deadRules.isEmpty()) {
            System.out.println("\nRules never triggered:");
            deadRules.stream().sorted().forEach(r -> System.out.println("  - " + r));
        }

        if (leastHitsLimit > 0 && !ruleHitCounts.isEmpty()) {
            int half = leastHitsLimit / 2;
            System.out.printf("\nHit distribution (top %d / bottom %d):\n", half, half);

            ruleHitCounts.entrySet().stream().sorted(Map.Entry.comparingByValue(reverseOrder())).limit(half).forEach(e -> System.out.printf("  - %-40s : %d hits\n", e.getKey(), e.getValue()));

            System.out.println("  ...");

            ruleHitCounts.entrySet().stream().sorted(Map.Entry.comparingByValue()).limit(half).sorted(Map.Entry.comparingByValue(reverseOrder())).forEach(e -> System.out.printf("  - %-40s : %d hits\n", e.getKey(), e.getValue()));
        }

        System.out.println("=".repeat(82));
    }
}