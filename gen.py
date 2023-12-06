import re

enhancedCodePattern = re.compile(r'\b\d+\.\d+\.\d+\b')

variables :'list[tuple[int,list[int] | None,str]]'= []

with open("raw.txt") as reader:
    for line in reader:
        if line.startswith('#'):
            continue

        parts = line.strip().split()

        code = int(parts[0])
        
        enhancedCode = None
        
        if enhancedCodePattern.fullmatch(parts[1]):
            enhancedCode = [int(num) for num in parts[1].split('.')]
        
        nameIndex = 2 if enhancedCode != None else 1

        variableName = "".join([ part.capitalize() for part in parts[nameIndex:] ]) 

        variables.append((code,enhancedCode,variableName))


with open("gen.go" ,'w') as writer:
    writer.write('package smtp_codes\n\n')
    for variable in variables:
        enhancedCode = 'NoEnhancedCodes'
        if variable[1] != None:
            enhancedCode = '[3]int16{' + ",".join([str(i) for i in variable[1]]) + '}'
        writer.write(f"var {variable[2]} = SmtpCode{{Code: {variable[0]}, EnhancedCodes: {enhancedCode }}}\n\n")




