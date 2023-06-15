import re

# Read the input file containing the output of the top command
with open('top_output.txt', 'r') as file:
    top_output = file.readlines()

# Regular expression pattern to match the lines containing process information
pattern = r'\s*(\d+)\s+.*\s+([a-zA-Z0-9_-]+)'

# Lists to store the extracted PIDs and users
pids = []
users = []

# Iterate over each line in the top output
for line in top_output:
    match = re.search(pattern, line)
    if match:
        # Extract PID and user from the matched line
        pid = match.group(1)
        user = match.group(2)

        # Store PID and user in separate lists
        pids.append(pid)
        users.append(user)

# Print the extracted PIDs and users
print("Running PIDs:")
for pid in pids:
    print(pid)

print("Users:")
for user in users:
    print(user)
