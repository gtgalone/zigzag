# Python 2.7

# Inject regexp module
import re

# Init a variable
foo = 'ab2v9bc13j5jf4jv21'

# Make a list with distinguished number in string after change type string to integer
bar = [int(s) for s in re.findall(r'\d+', foo)]

# Remove even in a list and then return number with square
def remove_even(list):
  return [pow(x, 2) for x in list if x % 2]
  
# Execute and put a variable
result = remove_even(bar)

# Print with sum of result list
print(sum(result))
