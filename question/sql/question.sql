-- 1 Please write a SQL statement to query the information of all courses where the number of students is between 50 and 55 in the course table courses.
SELECT * FROM courses
WHERE student_count >= 50 AND student_count <= 55
-- WHERE number_of_students BETWEEN 50 AND 55;

-- 2 Write a SQL statement to query the teachers table, teachers, to count the number of teachers of different ages, and return the results sorted by the age column, age, in descending order, with the column name of the counted number displayed as age_count.
SELECT age, COUNT(age) AS age_count
FROM teachers
GROUP BY age
ORDER BY age DESC;

-- 3 Write an SQL statement to delete all courses in the course table courses with a course creation date created_at before 2020.
DELETE FROM courses
WHERE created_at < '2020-01-01';

-- 4 Please write an SQL statement to update the email address of the teacher named Linghu Chong to "linghu.chong@lintcode.com" in the teacher table teachers.
UPDATE teachers
SET email = 'ddd@'  -- single quotes matter 
WHERE name = 'Linghu Chong';

-- 5 Write an SQL statement to change the number of students to 0 for all courses in the course table.
UPDATE courses
SET student_count = 0;


-- 6 Please write an SQL statement to query all the unique teacher's nationality (country) in the table teachers.
SELECT DISTINCT(country) FROM teachers; 


--  7 Write a SQL statement to make a left join between the courses table courses and the teachers table teachers, query the names of teachers with country = 'CN' and the names of courses they taught, with the result columns named course_name and teacher_name respectively.
SELECT c.name AS course_name, t.name AS teacher_name
FROM courses c
LEFT JOIN teachers t
ON c.teacher_id = t.id
WHERE t.country = 'CN';

-- --------------------------------------- ROLLBACK -----------------------------------------
-- 8 We need to undo the update of Linghu Chong's age in the teachers table. Please add the SQL statement to undo the update of Linghu Chong's age.
UPDATE teachers
SET age = 30
WHERE name = 'Linghu Chong';

-- ROLLBACK;
-- -----------------------------------------ROLLBACK ---------------------------------------

-- 9 Write an SQL statement to query the names of students who have the same name and return the results with the column name of the counted number displayed as name_count.
SELECT name, COUNT(*) AS name_count
FROM students
GROUP BY name
HAVING COUNT(*) > 1;

-- 10 Write a SQL statement to query the courses table to calculate the year difference between the time a course was created (created_at) and April 1, 2021, with the course name column displayed as courses_name, the time a course was created as courses_created_at, and the year difference column name displayed as year_diff.

SELECT name AS courses_name, created_at AS courses_created_at,
YEAR(created_at) - 2021 AS year_diff
FROM courses;
OR
SELECT 
    name AS courses_name,
    created_at AS courses_created_at,
    EXTRACT(YEAR FROM AGE('2021-04-01', created_at)) AS year_diff
FROM courses;
