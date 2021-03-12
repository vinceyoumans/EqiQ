# 10. Design and demonstrate, in code, an interface for a data access layer.


Scenario is that farmer has X fields and X.N cows on each field.
If the question is, how many cows does farmer have at anyone time, then
an external map would be constantly updated, and the calculation is done 
from a library... I am not entirely certian why I needed to use an interface, but I did it here.

All the Farmer needs to call in code is HowManyCowsDoIHave returning an int and err
