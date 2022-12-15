variable "input_file" {
  description = "Name of the input file"
  type        = string
  default     = "input.txt"
}

data "local_file" "input" {
  filename = "${path.module}/${var.input_file}"
}

locals {

  elfSplit                = split("\n\n", trim(data.local_file.input.content, "\n"))
  elfSplitSplit           = [for elf in local.elfSplit : split("\n", elf)]
  elfSplitSplitNum        = [for elf in local.elfSplitSplit : [for item in elf : tonumber(item)]]
  elfSplitSplitNumSum     = [for elf in local.elfSplitSplitNum : format("%09d", tostring(sum(elf)))]
  elfSplitSplitNumSumSort = reverse(sort(local.elfSplitSplitNumSum))
}

output "part1" {
  value = tonumber(local.elfSplitSplitNumSumSort[0])
}
output "part2" {
  value = local.elfSplitSplitNumSumSort[0] + local.elfSplitSplitNumSumSort[1] + local.elfSplitSplitNumSumSort[2]
}
