; ModuleID = 'feng_module'
source_filename = "feng_module"

declare i32 @printf(i8*, ...)

declare i32 @scanf(i8*, ...)

define i32 @main() {
entry:
  %a = alloca i32, align 4
  store i32 5, i32* %a, align 4
  %0 = load i32, i32* %a, align 4
  %1 = icmp sgt i32 %0, 1
  br i1 %1, label %7, label %4

2:                                                ; preds = %10, %7
  ret i32 1

3:                                                ; preds = %10, %4
  ret i32 0

4:                                                ; preds = %7, %entry
  %5 = load i32, i32* %a, align 4
  %6 = icmp sgt i32 %5, 3
  br i1 %6, label %10, label %3

7:                                                ; preds = %entry
  %8 = load i32, i32* %a, align 4
  %9 = icmp sgt i32 %8, 2
  br i1 %9, label %2, label %4

10:                                               ; preds = %4
  %11 = load i32, i32* %a, align 4
  %12 = icmp sgt i32 %11, 0
  br i1 %12, label %2, label %3
}
