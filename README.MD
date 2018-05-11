SQL
TABLE
  Clients
    id, first name, last name, middlename
    CREATE TABLE `clients` (
        `id` INT NOT NULL AUTO_INCREMENT,
        `firstname` VARCHAR(64) NULL DEFAULT NULL,
        `lastname` VARCHAR(64) NULL DEFAULT NULL,
        `middlename` VARCHAR(64) NULL DEFAULT NULL,
        PRIMARY KEY (`id`)
    );
  Services
    name, groupname`
    CREATE TABLE `services` (
        `name` VARCHAR(64) NULL DEFAULT NULL,
        `groupname` VARCHAR(64) NULL DEFAULT NULL,
        PRIMARY KEY (`name`)
    );
  TimeSheets
    clientid, day of week, date, weekstartdate, service, employee

HTML
select year, week, (default to current week/year) client
Features
  copy from day, copy from week, edit day, print

# Checkout https://github.com/jung-kurt/gofpdf for PDF printing the timesheets once completed
