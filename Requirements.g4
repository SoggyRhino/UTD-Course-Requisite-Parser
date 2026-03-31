grammar Requirements;


prog     : sentence+ EOF ;
sentence : requisite PERIOD? ;

requisite
    : requisite AND EXCLUDE_DMHP_LLC_NOTICE          # excludeNoticeReq
    | requisite AND academic_plan_condition          # appendAcademicPlanReq
    | academic_plan_condition                        # academicPlanReq
    | EXACT_COREQ_NOTICE                             # exactCoreqNoticeReq
    | COMPUTER_SCHOLARS_PROGRAM                      # computerScholarsReq
    | gpa_repeate_rule                               # gpaRepeatReq
    | repeat_limit_hours_rule                        # repeatLimitHoursReq
    | repeat_limit_times_rule                        # repeatLimitTimesReq
    | repeat_rule                                    # repeatReq
    | degree_satisfaction_rule                       # degreeSatisfactionReq
    | credit_for_rule                                # creditForReq
    | living_learning_rule                           # livingLearningReq
    | school_rule                                    # schoolReq
    | major_condition                                # majorReq
    | PREREQ_KW COLON? expr (OR EQUIVALENT)?         # prereqReq
    | COREQ_KW  COLON? expr (OR EQUIVALENT)?         # coreqReq
    | PREREQ_KW COLON? expr AND COREQ_KW COLON expr  # prereqAndCoreqReq
    | PRE_OR_CO_KW COLON? expr (OR EQUIVALENT)?      # preOrCoReq
    | '(' SAME_AS expr ')'                           # sameAsParenReq
    | SAME_AS expr                                   # sameAsReq
    | expr                                           # exprReq
    ;

expr
    : '(' expr ')'                          # parenExpr
    | expr COMMA? OR expr                   # orExpr
    | expr COMMA? AND expr                  # andExpr
    | expr COMMA? AMPERSAND expr            # ampersandExpr
    | expr OR EQUIVALENT                    # equivalentExpr
    | INTRUCTOR_CONSENT                     # instructorConsentExpr
    | DEPARTMENT_CONSENT                    # departmentConsentExpr
    | UTEACH_CONSENT                        # uteachConsentExpr
    | UPPER_DVISION_STANDING                # upperDivisionStandingExpr
    | COMPUTER_SCHOLARS_PROGRAM             # computerScholarsExpr
    | GOOD_ACADEMIC_STANDING                # goodAcademicStandingExpr
    | gpa_condition                         # gpaExpr
    | group_condition                       # groupExpr
    | concurrent_enrollment_condition       # concurrentEnrollmentExpr
    | grade_condition                       # gradeExpr
    | alternative_condition                 # alternativeExpr
    | grade_level_standing_condition        # gradeLevelStandingExpr
    | graduate_standing_condition           # graduateStandingExpr
    | major_condition                       # majorExpr
    | degree_condition                      # degreeExpr
    | core_condition                        # coreExpr
    | any_core_condition                    # anyCoreExpr
    | complete_n_condition                  # completeNExpr
    | semester_credit_hours_condition       # semesterCreditHoursExpr
    | minimum_hours_condition               # minimumHoursExpr
    | upper_division_hours_condition        # upperDivisionHoursExpr
    | uppper_division_classes_condition     # upperDivisionClassesExpr
    | research_condition                    # researchExpr
    | placement_test_condition              # placementTestExpr
    | ap_score_condition                    # apScoreExpr
    | aleks_score_condition                 # aleksScoreExpr
    | exact_section_condition               # exactSectionExpr
    | any_major_course_condition            # anyMajorCourseExpr
    | living_learning_rule                  # livingLearningExpr
    | repeat_rule                           # repeatRuleExpr
    | repeat_limit_hours_rule               # repeatLimitHoursExpr
    | course                                # courseExpr
    ;


// Course primitives
course
    : '(' course ')'                        # parenCourse
    | course '/' course                     # crossListedCourse
    | PREFIX ('/' PREFIX)* COURSE_NUMBER    # simpleCourse
    ;

course_list
    : course (OR course)+                           # fullCourseList
    | PREFIX COURSE_NUMBER (OR COURSE_NUMBER)+      # shorthandCourseList
    ;

title : (CAPITALIZED | CORE)+ (AND (CAPITALIZED | CORE)+)* ;

degree_atom : CAPITALIZED | CORE | WORD ;

degree : degree_atom+ (AND degree_atom+)* ;

degree_list : degree ((COMMA AND | COMMA | OR) degree)* ;

// Grade conditions
grade_condition
    : course WITH_GRADE GRADE OR_BETTER?                                   # simpleGradeCondition
    | course WITH_GRADE 'greater than or equal to' GRADE ('(' GPA ')')?    # gpaGradeCondition
    | grade_course_list WITH_GRADE GRADE OR_BETTER?                        # gradeListCondition
    | A_GRADE_OF_AT_LEAST GRADE OR_BETTER? 'in' grade_course_list          # gradeAtLeastCondition
    ;

grade_course_list
    : 'either'? course (OR 'in'? course)*             #eitherGradeCourseList
    |  course (AND course)*                           #allGradeCourseList
    | '(' course (OR course)* ')'                     #parenGradeCourseList
    ;

alternative_condition
    : course OR EQUIVALENT                    # courseAlternativeCondition
    | course_list OR EQUIVALENT               # gradeCourseListAlternativeCondition
    ;


// Standing conditions
grade_level_standing_condition
    : GRADE_LEVEL (OR GRADE_LEVEL)* LEVEL? STANDING                                        # gradeLevelStandingCondition
    | GRADE_LEVEL degree MAJOR_KW STANDING                                                  # gradeLevelMajorStandingCondition
    | MINIMUM_OF GRADE_LEVEL STANDING                                                       # minimumGradeLevelStandingCondition
    | AT_LEAST GRADE_LEVEL (DASH LEVEL | LEVEL)? STANDING                                  # atLeastGradeLevelStandingCondition
    | PREFIX MAJOR_KW? ONLY_KW 'with' GRADE_LEVEL (OR GRADE_LEVEL)* LEVEL? STANDING        # prefixGradeLevelStandingCondition
    ;

graduate_standing_condition
    : 'Graduate standing in' degree    # graduateStandingInCondition
    | 'Graduate Level Standing'        # graduateLevelStandingCondition
    ;


// GPA conditions
gpa_condition
    : UNIVERSITY_GPA_KW GPA                                                 # universityGpaCondition
    | 'Minimum GPA requirement' GPA                                         # minimumGpaCondition
    | 'a GPA of' GPA OR_BETTER 'in' degree COURSE_KW?                       # gpaInCourseCondition
    ;


// Major / degree conditions
major_condition
    : PREFIX (OR PREFIX)* (DIVISION_TYPE | DEGREE_LEVEL)? GRADE_LEVEL? MAJOR_KW? ONLY_KW?     # prefixMajorCondition
    | GRADE_LEVEL PREFIX MAJOR_KW? ONLY_KW?                                                  # gradeLevelPrefixMajorCondition
    | DEGREE_LEVEL PREFIX MAJOR_KW? ONLY_KW?                                                  # degreeTypePrefixMajorCondition
    | degree MAJOR_KW ONLY_KW?                                                               # namedMajorCondition
    | degree DEGREE_LEVEL MAJOR_KW? ONLY_KW?                                                  # namedDegreeTypeMajorCondition
    ;

degree_condition
    : 'an undergraduate degree in' degree 'and adequate foundation/academic performance in a corresponding area'  # undergraduateDegreeCondition
    | 'Bachelor\'s or Master\'s degree in' degree_list                                                            # bachelorsOrMastersCondition
    ;


// Core conditions
core_condition : COMPLETION_OF ('a' | 'an')? CORE_NUMBER (CORE | '(' CORE ')' )? CORE_KW COURSE_KW?   # coreCondition;

any_core_condition : 'any' SMALL_INT SEMESTER_CREDIT_HOURS CORE_NUMBER CORE_KW COURSE_KW   # anyCoreSCHCondition;


// Hours / credit conditions

semester_credit_hours_condition
    : INT SEMESTER_CREDIT_HOURS                                              # semesterCreditHoursCondition
    ;

minimum_hours_condition
    : MINIMUM_OF SMALL_INT SEMESTER_CREDIT_HOURS 'in any combination of' course_list   # minimumHoursCondition
    | AT_LEAST SMALL_INT 'semester credits of' course_list                             # minimumHoursOfCondition
    | SMALL_INT SEMESTER_CREDIT_HOURS 'from the following' COLON course_list           # minimumHoursFromCondition
    ;

upper_division_hours_condition
    : SMALL_INT 'SCH of upper-division' PREFIX COURSE_KW        # upperDivisionSCHCondition
    ;

uppper_division_classes_condition
    : AT_LEAST NUMBER_STRING PREFIX UPPER_DIVISION_COURSE_NUMBER COURSE_KW   # upperDivisionCountCondition
    | 'a 4000-level' PREFIX COURSE_KW                                        # upperDivisionSingleCondition
    ;

research_condition
    : AT_LEAST SMALL_INT SEMESTER_CREDIT_HOURS 'of' DIVISION_TYPE 'research'   # researchCondition
    ;

complete_n_condition
    : COMPLETION_OF NUMBER_STRING 'of the following' COLON course_list       # completeNOfFollowingCondition
    | NUMBER_STRING COURSE_KW 'from the following' DASH course_list          # completeNFromFollowingCondition
    ;


// Placement / score conditions

placement_test_condition
    : 'a'? placement_test_name SCORE_KW (LESS_THAN | GREATER_THAN) INT    # placementScoreComparisonCondition
    | 'a'? placement_test_name SCORE_KW 'of' INT DASH INT                 # placementScoreRangeCondition
    | 'a'? placement_test_name 'of'? INT OR_BETTER                        # placementScoreMinimumCondition
    ;

placement_test_name
    : PREFIX PLACEMENT_KW TEST_KW
    | title PLACEMENT_KW TEST_KW?
    ;

ap_score_condition
    : 'AP score of' AT_LEAST SMALL_INT   # apScoreCondition
    ;

aleks_score_condition
    : A_SORCE_OF INT '%' 'on ALEKS math placement exam'   # aleksScoreCondition
    ;


// People / group conditions

group_condition
    : 'Students in both' group '/' group STUDENT_GROUP ONLY_KW                              # bothGroupCondition
    | group (AND group)* (STUDENT_GROUP | STUDENTS) ONLY_KW?                                # groupListCondition
    ;

group
    : COLLEGIUM_V_HONORS
    | LIBERAL_ARTS_HONORS
    | SCVG
    | DMHP
    | DLAH
    ;


// Misc conditions
concurrent_enrollment_condition
    : CONCURRENT_ENROLLMENT_IN course   # concurrentEnrollmentCondition
    ;

exact_section_condition
    : course PERIOD SECTION_NUMBER     # exactSectionCondition
    | course 'workshop' SECTION_NUMBER # workshopSectionCondition
    ;

any_major_course_condition
    : ANY_PREVIOUS PREFIX COURSE_KW   # anyPreviousMajorCourseCondition
    ;

academic_plan_condition
    : 'Academic Plan' (NOT_EQUAL | EQUAL) 'to' ACADEMIC_PLAN # academicPlanCondition
    ;


// Rules

repeat_rule
    : course REPEAT_RESTRICTION              # courseRepeatRule
    | PREFIX 'Internship' REPEAT_RESTRICTION # internshipRepeatRule
    | REPEAT_RESTRICTION                     # bareRepeatRule
    ;

repeat_limit_hours_rule
    : REPEAT_LIMIT DASH course ONLY_KW? 'may' ONLY_KW? 'be repeated' ONLY_KW? 'for'? 'a maximum of' SMALL_INT SEMESTER_CREDIT_HOURS                    # repeatMaxHoursRule
    | REPEAT_LIMIT DASH course ONLY_KW? 'may' ONLY_KW? 'be repeated' ONLY_KW? 'for'? SMALL_INT SEMESTER_CREDIT_HOURS 'maximum'                         # repeatHoursMaxSuffixRule
    | course REPEAT_LIMIT DASH 'This' COURSE_KW ('may' | 'can') ONLY_KW? 'be repeated' ONLY_KW? 'for'? 'a maximum of' SMALL_INT SEMESTER_CREDIT_HOURS  # courseRepeatMaxHoursRule
    | REPEAT_LIMIT DASH course AND course 'combined may only be repeated for a maximum of' SMALL_INT SEMESTER_CREDIT_HOURS                             # combinedRepeatMaxHoursRule
    | course REPEAT_LIMIT DASH 'May be repeated for credit as topics vary' SMALL_INT SEMESTER_CREDIT_HOURS 'maximum'                                   # topicsVaryRepeatRule
    | course REPEAT_LIMIT                                                                                                                              # courseRepeatLimitRule
    ;

repeat_limit_times_rule
    : REPEAT_LIMIT DASH course 'may' ONLY_KW? 'be repeated up to' SMALL_INT 'times'          # repeatUpToTimesRule
    | REPEAT_LIMIT DASH course 'may' ONLY_KW? 'be repeated' 'a maximum of' SMALL_INT 'times' # repeatMaxTimesRule
    ;

// Im not sure what this means but this is what gemini says it means
// students not enrolled in "" program face a rule if they try
// to retake course to improve their GPA.
gpa_repeate_rule
    : 'GPA Repeat Restriction' DASH course  # gpaRepeatRule
    ;

degree_satisfaction_rule
    : MAY_NOT_BE_USED_TO_SATISFY DEGREE_LEVEL? PREFIX 'degree requirements'                                          # prefixDegreeSatisfactionRule
    | MAY_NOT_BE_USED_TO_SATISFY 'degree requirements in' DEGREE_LEVEL? title                                        # namedDegreeSatisfactionRule
    | MAY_NOT_BE_USED_TO_SATISFY 'degree requirements' (OR? THE? DEGREE_LEVEL? PREFIX)+ 'degree plans'               # multiPrefixDegreeSatisfactionRule
    | MAY_NOT_BE_USED_TO_SATISFY 'degree requirements for' (OR? THE? DEGREE_LEVEL? PREFIX)+ 'degree plans'           # multiPrefixForDegreeSatisfactionRule
    | MAY_NOT_BE_USED_TO_SATISFY 'the degree requirements of' (OR? THE? DEGREE_LEVEL? PREFIX)+ 'degree plans'        # ofMultiPrefixDegreeSatisfactionRule
    | MAY_NOT_BE_USED_TO_SATISFY 'degree requirements for' DEGREE_LEVEL? 'majors in'? 'the School of'? degree+       # schoolDegreeSatisfactionRule
    | MAY_NOT_BE_USED_TO_SATISFY 'degree requirements for' DEGREE_LEVEL? 'majors in'? 'Schools of'? degree_list+     # schoolsDegreeSatisfactionRule
    | MAY_NOT_BE_USED_TO_SATISFY 'degree requirements by students in' degree                                         # studentDegreeSatisfactionRule
    | MAY_NOT_BE_USED_TO_SATISFY 'mathematics requirements by students in Mathematics'                               # mathDegreeSatisfactionRule
    | degree_satisfaction_rule AND 'may not be used to satisfy electives'                                            # electivesDegreeSatisfactionRule
    ;

credit_for_rule
    : CREDIT_PREFIX (COLON|COMMA)? expr
    ;

living_learning_rule
    : PREFIX ('&' PREFIX)* LIVING_LEARNING_COMMUNITY    # prefixLivingLearningRule
    | degree_list LIVING_LEARNING_COMMUNITY             # namedLivingLearningRule
    ;

school_rule
    : 'Open to students in the School of' degree_list ONLY_KW?
    ;

// Lexer

// Long literal notices (must come before any keywords they contain)
EXCLUDE_DMHP_LLC_NOTICE : 'non-DMHP/non-LLC (DMLC, DFLC, DHLC) Student Group Only' ;
EXACT_COREQ_NOTICE      : 'Check class notes to make sure you are selecting the matching corequisite section';
CONCURRENT_ENROLLMENT_IN: 'concurrent enrollment in';
GOOD_ACADEMIC_STANDING  : 'good academic standing';
CREDIT_PREFIX           : 'Credit cannot be received for more than one of the following'
                        | 'Credit cannot be received for both courses'
                        | 'Credit cannot be received for both'
                        ;
MAY_NOT_BE_USED_TO_SATISFY : 'May not be used to satisfy' | 'May not be used to fulfill';
UNIVERSITY_GPA_KW          : ('a' | 'A') ' university grade point average' (' (GPA)')? ' of at least';


// Requisite keywords
PRE_OR_CO_KW : 'Prerequisite or Corequisite'
             | 'Prerequisites or Corequisites'
             | 'Corequisite or Prerequisite'
             | 'Corequisites or Prerequisites'
             ;
PREREQ_KW : 'Prerequisite' 's'?;
COREQ_KW  : 'Corequisite'  's'? | 'Coreqrequisite';


// Grade / score keywords
WITH_GRADE : 'with a minimum grade of'
           | 'with a grade of'
           | 'with a grade'
           | 'with the grade'
           | 'with grade'
           | 'with a'
           | 'with'
           ;

A_GRADE_OF_AT_LEAST : 'A grade of at least a'
                    | 'a grade of at least a'
                    | 'A grade of'
                    ;

A_SORCE_OF          : 'A score of'
                    | 'a score of'
                    | 'A minimal placement score of'
                    | 'a minimal placement score of'
                    ;

OR_BETTER    : 'or better' | 'or higher';
SCORE_KW     : [Ss]'core';
PLACEMENT_KW : [Pp]'lacement' | 'Place';
TEST_KW      : [Tt]'est';


// Standing / level keywords
UPPER_DVISION_STANDING : 'Upper-division standing';
GRADE_LEVEL : [Ff]'reshm'[ae]'n'
            | [Ss]'ophomore'
            | [Jj]'unior' | 'JR'
            | [Ss]'enior' | 'SR'
            ;
STANDING  : [Ss]'tanding';
LEVEL     : [Ll]'evel';
EQUIVALENT: [Ee]'quivalent';


// Major / degree keywords


DIVISION_TYPE : [Uu]'ndergraduate' | 'Ugrd' | 'ugrd'
              | [Gg]'raduate'      | 'GRAD' | 'grad'
              | [Dd]'octoral'
              ;
DEGREE_LEVEL : 'MS' | 'BS' | 'PHD' | 'PhD';
MAJOR_KW    : [Mm]'ajor''s'?;
ONLY_KW     : [Oo]'nly';
STUDENTS    : [Ss]'tudents';


// Core / course keywords
CORE : 'Communication'
     | 'Mathematics'
     | 'Life and Physical Sciences'
     | 'Language, Philosophy and Culture'
     | 'Creative Arts'
     | 'American History'
     | 'Government/Political Science'
     | 'Social and Behavioral Sciences'
     | 'Component Area Option'
     ;


CORE_KW   : [Cc]'ore';
COURSE_KW : [Cc]'lass''es'? | [Cc]'ourse''s'? | 'coursework';

COMPLETION_OF       : [Cc]'ompletion of';
SEMESTER_CREDIT_HOURS: 'semester credit hour''s'?;
MINIMUM_OF          : [Mm]'inimum of';
AT_LEAST            : [Aa]'t least' | [Aa]'t Least';
LIVING_LEARNING_COMMUNITY: 'Living Learning Community';
COMPUTER_SCHOLARS_PROGRAM: 'Computer Scholars Program' | 'Computing Scholars Program';
ANY_PREVIOUS        : 'any previous' | 'Any previous';


// Honors / group tokens
COLLEGIUM_V_HONORS : 'Collegium V Honors' | 'CV Honors';
LIBERAL_ARTS_HONORS: 'Liberal Arts Honors';
STUDENT_GROUP      : 'Student Group''s'? | 'student group''s'? | [Gg]'roup';
SCVG : 'SCVG';
DMHP : 'DMHP';
DLAH : 'DLAH';


// Consent tokens
INTRUCTOR_CONSENT  : 'instructor consent' | 'instructor consent required';
DEPARTMENT_CONSENT : 'department consent' | 'department consent required';
UTEACH_CONSENT     : 'UTeach advisor consent required';


// Repeat tokens
REPEAT_RESTRICTION : 'Repeat Restricition' | 'Repeat Restriction' | 'Repeat Restiction';
REPEAT_LIMIT       : 'Repeat Limit';


// Operators and punctuation
AND       : 'and';
OR        : 'or';
COMMA     : ',';
AMPERSAND : '&';
SAME_AS   : 'Same as';
EQUAL     : 'Equal';
NOT_EQUAL : 'Not Equal';
LESS_THAN    : 'less than';
GREATER_THAN : 'greater than';

DASH   : '-';
COLON  : ':';
PERIOD : '.';

FOR        : 'for';
THE        : 'the' | 'The';
OF         : 'of';
MAY_KW     : 'may';
BE_REPEATED: 'be repeated';


// Numeric tokens

NUMBER_STRING : [Oo]'ne' | [Tt]'wo'   | [Tt]'hree' | [Ff]'our' | [Ff]'ive'
              | [Ss]'ix' | [Ss]'even' | [Ee]'ight' | [Nn]'ine'
              ;


CORE_NUMBER                : [0][1-6][0];
SECTION_NUMBER             : [015][0-9][0-9];
SMALL_INT : [1-9] | '1'[0-9];          // 1-19
INT       : '0'
          | [2-9][0-9]                  // 20-99
          | [1-9][0-9][0-9]             // 100-999
          ;
GPA : [0-4][.][0-9][0-9]?[0-9]?;

// Structural / pattern tokens

ACADEMIC_PLAN              : [A-Z][A-Z][A-Z][A-Z][A-Z][A-Z]+;
PREFIX                     : [A-Z][A-Z][A-Z]?[A-Z]? | 'Math';
COURSE_NUMBER              : [0-9][0-9vV][0-9][0-9];
UPPER_DIVISION_COURSE_NUMBER: [4][1-9Xx][1-9Xx][1-9Xx];
GRADE                      : [A-F][+-]?;

// Generic tokens (must come last)

CAPITALIZED : [A-Z][a-z]+;
WORD        : [a-zA-Z]+;

WS    : [ \t\n\r]+ -> skip;

ANY_OTHER : . -> skip;