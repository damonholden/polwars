// Assumptions:
// - the two parameters passed into the function will always be a string type with the format `YYYY-MM-DD hh:mm:ss`

// Limitations:
// - can handle timescales in the same hour, same day, or over two days in the same month
// - can currently only handle a single pair of dates (i.e. a start and finish times)
// - currently needs the window object of the browser context to generate a csv, so this will not work in node, etc.

const timeKeeper = (timeA, timeB) => {
  const timeSplit = time => {
    const date = time.substring(0, 10);
    const hours = parseInt(time.substring(11, 13));
    const minutes = parseInt(time.substring(14, 16));
    const seconds = parseInt(time.substring(17, 19));
    const day = parseInt(date.substring(8, 10));

    return { date, day, hours, minutes, seconds };
  };

  const start = timeSplit(timeA);
  const finish = timeSplit(timeB);

  const hours = [
    0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
    21, 22, 23,
  ];

  const contentRow = [];

  if (start.date === finish.date) {
    // As both dates passed into the function are the same, there will only be one line in the csv file.

    const onlyRow = [];

    for (const hour of hours) {
      if (hour < start.hours) {
        onlyRow.push('0');
      } else if (hour === start.hours) {
        if (start.hours === finish.hours) {
          onlyRow.push(
            60 * (finish.minutes - start.minutes) +
              (finish.seconds - start.seconds)
          );
        } else {
          onlyRow.push(60 * (60 - start.minutes) - start.seconds);
        }
      } else if (hour < finish.hours) {
        onlyRow.push('3600');
      } else if (hour === finish.hours) {
        onlyRow.push(60 * finish.minutes + finish.seconds);
      } else onlyRow.push('0');
    }

    contentRow.push(onlyRow);
  } else {
    // As both dates are not the same, the csv file will contain multiple lines.

    for (let i = start.day; i <= finish.day; i++) {
      let row = [];
      if (i === start.day) {
        for (const hour of hours) {
          if (hour < start.hours) {
            row.push('0');
          } else if (hour === start.hours) {
            row.push(60 * (60 - start.minutes) - start.seconds);
          } else {
            row.push(`3600`);
          }
        }
      } else if (i < finish.day && i !== finish.day) {
        for (const hour of hours) {
          row.push(`3600`);
        }
      } else if (i === finish.day) {
        for (const hour of hours) {
          if (hour < finish.hours) {
            row.push('3600');
          } else if (hour === finish.hours) {
            row.push(60 * finish.minutes + finish.seconds);
          } else {
            row.push(`0`);
          }
        }
      }
      contentRow.push(row);
    }
  }

  const rows = [
    [`DATE`, ...hours.map(hour => `HOUR_${hour}`)],
    ...contentRow.map((row, index, array) => {
      if (index + 1 === array.length) {
        return [finish.date, ...row];
      } else {
        return [start.date, ...row];
      }
    }),
  ];

  const csvContent =
    'data:text/csv;charset=utf-8,' + rows.map(e => e.join(',')).join('\n');

  const encodedUri = encodeURI(csvContent);

  window.open(encodedUri);
};

// TODO: use string literals to add date strings into large number to support larger time spans over different months and years
// TODO: refactor codebase to make it less ugly
