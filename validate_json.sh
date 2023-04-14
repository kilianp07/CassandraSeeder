#!/bin/bash

# Input file path
input_file="restaurants.json"

# Output file path
output_file="restaurant2.json"

# Get the total number of lines in the input file
total_lines=$(wc -l < "$input_file")

# Write opening square bracket to the output file
echo "[" > "$output_file"

# Loop through each line in the input file
while IFS= read -r line; do
  # Check if this is the last line
  if [[ $((++line_count)) -eq $total_lines ]]; then
    # If it's the last line, remove the trailing comma
    echo "$line" >> "$output_file"
  else
    # If it's not the last line, append a comma at the end of the line
    echo "$line," >> "$output_file"
  fi
done < "$input_file"

# Write closing square bracket to the output file
echo "]" >> "$output_file"

echo "Commas added to the end of each line (except the last line), and square brackets added to the beginning and end of the file. Output written to $output_file"
