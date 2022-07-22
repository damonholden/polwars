const timeKeeper = (timeA, timeB) => {
  const timeSplit = time => {
    const date = time.substring(0, 10);
    const hours = time.substring(11, 13);
    const minutes = time.substring(14, 16);
    const seconds = time.substring(17, 19);

    return { date, hours, minutes, seconds };
  };

  const clockIn = timeSplit(timeA);
  const clockOut = timeSplit(timeB);

  const headerRow = [
    `HOUR_0`,
    `HOUR_1`,
    `HOUR_2`,
    `HOUR_3`,
    `HOUR_4`,
    `HOUR_5`,
    `HOUR_6`,
    `HOUR_7`,
    `HOUR_8`,
    `HOUR_9`,
    `HOUR_10`,
    `HOUR_11`,
    `HOUR_12`,
    `HOUR_13`,
    `HOUR_14`,
    `HOUR_15`,
    `HOUR_16`,
    `HOUR_17`,
    `HOUR_18`,
    `HOUR_19`,
    `HOUR_20`,
    `HOUR_21`,
    `HOUR_22`,
    `HOUR_23`,
  ];

  const testRow = [
    `2021 - 12 - 24`,
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    1459,
    3600,
    3600,
    3600,
    3600,
    3600,
    3600,
    3600,
    3600,
    3230,
    0,
    0,
    0,
    0,
    0,
    0,
  ];

  const contentRow = [];

  for (header of headerRow) {
    if (parseInt(header.replace(`HOUR_`, ``)) < parseInt(clockIn.hours)) {
      contentRow.push('0');
    } else if (
      parseInt(header.replace(`HOUR_`, ``)) === parseInt(clockIn.hours)
    ) {
      contentRow.push(
        (
          60 * (60 - parseInt(clockIn.minutes)) -
          parseInt(clockIn.seconds)
        ).toString()
      );
    } else if (
      parseInt(header.replace(`HOUR_`, ``)) < parseInt(clockOut.hours)
    ) {
      contentRow.push('3600');
    } else if (
      parseInt(header.replace(`HOUR_`, ``)) === parseInt(clockOut.hours)
    ) {
      contentRow.push(
        (
          60 * parseInt(clockOut.minutes) +
          parseInt(clockOut.seconds)
        ).toString()
      );
    } else contentRow.push('0');
  }

  const rows = [
    [`DATE`, ...headerRow],
    [testRow],
    [clockIn.date, ...contentRow],
  ];

  let csvContent =
    'data:text/csv;charset=utf-8,' + rows.map(e => e.join(',')).join('\n');

  var encodedUri = encodeURI(csvContent);
  window.open(encodedUri);
};

timeKeeper(`2021-12-24 08:35:41`, `2021-12-24 17:53:50`);
