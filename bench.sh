#!/bin/bash

# Detect if output is to a terminal
if [ -t 1 ]; then
  IS_TERMINAL=1
  USE_COLOR=1
else
  IS_TERMINAL=0
  USE_COLOR=0
fi

# Run benchmark
if [ -z "$1" ]; then
  if [ "$IS_TERMINAL" = "1" ]; then
    # Stream to terminal and capture
    output=$(go test -bench=BenchmarkAll -benchmem -benchtime=1s 2>&1 | tee /dev/stderr)
  else
    # Just capture, print everything at the end
    output=$(go test -bench=BenchmarkAll -benchmem -benchtime=1s 2>&1)
    echo "$output"
  fi
else
  if [ "$IS_TERMINAL" = "1" ]; then
    # Stream to terminal and capture
    output=$(go test -bench=BenchmarkAll/day${1}_ -benchmem -benchtime=1s 2>&1 | tee /dev/stderr)
  else
    # Just capture, print everything at the end
    output=$(go test -bench=BenchmarkAll/day${1}_ -benchmem -benchtime=1s 2>&1)
    echo "$output"
  fi
fi

# Parse and format the benchmark results
echo ""
echo "=== Human Readable ==="
echo "$output" | awk -v use_color="$USE_COLOR" '
/^Benchmark/ {
    # ANSI color codes (only if terminal)
    if (use_color == 1) {
        RED = "\033[31m"
        RESET = "\033[0m"
    } else {
        RED = ""
        RESET = ""
    }
    
    # Extract values
    name = $1
    iters = $2
    ns = $3
    bytes = $5
    allocs = $7
    
    # Convert time
    if (ns >= 1000000000) {
        time = sprintf("%.2fs", ns/1000000000)
    } else if (ns >= 1000000) {
        time = sprintf("%.2fms", ns/1000000)
    } else if (ns >= 1000) {
        time = sprintf("%.2fµs", ns/1000)
    } else {
        time = sprintf("%.0fns", ns)
    }
    
    # Convert memory
    if (bytes >= 1073741824) {
        mem = sprintf("%.2fGB", bytes/1073741824)
    } else if (bytes >= 1048576) {
        mem = sprintf("%.2fMB", bytes/1048576)
    } else if (bytes >= 1024) {
        mem = sprintf("%.2fKB", bytes/1024)
    } else {
        mem = sprintf("%.0fB", bytes)
    }
    
    # Format allocations
    if (allocs >= 1000000) {
        alloc_str = sprintf("%.2fM", allocs/1000000)
    } else if (allocs >= 1000) {
        alloc_str = sprintf("%.2fK", allocs/1000)
    } else {
        alloc_str = sprintf("%.0f", allocs)
    }
    
    # Extract just the day/part from name
    split(name, parts, "/")
    if (length(parts) > 1) {
        short_name = parts[2]
    } else {
        short_name = name
    }
    
    # Determine if we need color (track for padding adjustment)
    time_colored = 0
    mem_colored = 0
    
    # Color time if > 100ms (100,000,000 ns) and color is enabled
    if (ns >= 100000000 && use_color == 1) {
        time = RED time RESET
        time_colored = 1
    }
    
    # Color memory if > 10MB (10,485,760 bytes) and color is enabled
    if (bytes >= 10485760 && use_color == 1) {
        mem = RED mem RESET
        mem_colored = 1
    }
    
    # Adjust format width for color codes (RED + RESET = 9 chars) only if using color
    time_width = (time_colored && use_color == 1) ? 17 : 8
    mem_width = (mem_colored && use_color == 1) ? 19 : 10
    
    printf "%-20s %" time_width "s  %" mem_width "s  %8s allocs\n", short_name, time, mem, alloc_str
}
'
