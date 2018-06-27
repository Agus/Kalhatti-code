import json
import pip
import sys
import time

def install(package):
    try:
        result = pip.main(['install','-q', package])
        return result == 0;
    except Exception as e:
        return false;


# Main section

with open('dependencies.json') as f:
    data = json.load(f)
failedDependencies = [];
total = len(data["Dependencies"]);
for i, dependency in enumerate(data["Dependencies"]):
    sys.stdout.write("Install progress: %d of %d   \r" % (i+1,total) )
    sys.stdout.flush()
    if install(dependency) == False:
        failedDependencies.append(dependency)
print('')
if len(failedDependencies) > 0:
    print('-----Failed to install the following packages-----')
    print('\r\n'.join(failedDependencies))
