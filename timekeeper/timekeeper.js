const timeKeeper = (timeA, timeB) => {
  const timeSplit = time => {
    const date = time.substring(0, 10);
    const hours = parseInt(time.substring(11, 13));
    const minutes = parseInt(time.substring(14, 16));
    const seconds = parseInt(time.substring(17, 19));

    return { date, hours, minutes, seconds };
  };

  const clockIn = timeSplit(timeA);
  const clockOut = timeSplit(timeB);

  const hours = [
    0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
    21, 22, 23,
  ];

  const contentRow = [];

  for (header of hours) {
    if (header < clockIn.hours) {
      contentRow.push('0');
    } else if (header === clockIn.hours) {
      if (clockIn.hours === clockOut.hours) {
        contentRow.push(
          60 * (clockOut.minutes - clockIn.minutes) +
            (clockOut.seconds - clockIn.seconds)
        );
      } else {
        contentRow.push(60 * (60 - clockIn.minutes) - clockIn.seconds);
      }
    } else if (header < clockOut.hours) {
      contentRow.push('3600');
    } else if (header === clockOut.hours) {
      contentRow.push(60 * clockOut.minutes + clockOut.seconds);
    } else contentRow.push('0');
  }

  const rows = [
    [`DATE`, ...hours.map(header => `HOUR_${header}`)],
    [clockIn.date, ...contentRow],
  ];

  const csvContent =
    'data:text/csv;charset=utf-8,' + rows.map(e => e.join(',')).join('\n');

  const encodedUri = encodeURI(csvContent);

  window.open(encodedUri);
};

timeKeeper(`2021-12-24 08:35:41`, `2021-12-24 08:45:51`);
