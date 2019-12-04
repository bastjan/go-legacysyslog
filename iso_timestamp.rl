// Original work Copyright (c) 2018, InfluxData Inc.

%%{
machine iso_timestamp;

# 0..59
sexagesimal = '0'..'5' . '0'..'9';

# 01..31
datemday = ('0' . '1'..'9' | '1'..'2' . '0'..'9' | '3' . '0'..'1');

# 01..12
datemonth = ('0' . '1'..'9' | '1' . '0'..'2');

datefullyear = digit{4};

fulldate = datefullyear '-' datemonth '-' datemday;

# 01..23
timehour = ('0'..'1' . '0'..'9' | '2' . '0'..'3');

timeminute = sexagesimal;

timesecond = sexagesimal;

timesecfrac = '.' digit{1,6};

timenumoffset = ('+' | '-') timehour ':' timeminute;

timeoffset = 'Z' | timenumoffset;

partialtime = timehour ':' timeminute ':' timesecond . timesecfrac?;

fulltime = partialtime . timeoffset;

iso_timestamp = fulldate 'T' fulltime;

}%%
