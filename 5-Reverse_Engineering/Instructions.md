# Reverse Engineering
Go is extremely easy to reverse engineer because it doesn't natively come with the ability to obfuscate.

All function names are in cleartext stored within custom tables alongside the standard PE or ELF header sections.

All strings are referenced by tables and most strings are held as a continguous block of text. As in, all strings in Go are listed back to back, with no spacing or nulls between strings. Grabbing a string in this "block" is done by a pointer to the start of a string paired with the length.

We will be performing the reverse engineering of Go in Ghidra.

## Labeling Functions
Download this, build it and run it against a Go binary. Store the result in a file.

https://github.com/mandiant/GoReSym

Then go into script manager and make a new script, just make this the contents. 
https://github.com/mandiant/GoReSym/blob/master/GhidraPython/goresym_rename.py

Place this line at the top of the script. This lets Ghidra know it is using Python2.x to run this script.
```
# @runtime Jython
```

If it doesn't load because you don't have the extension, go back to the project window of Ghidra, ```File -> Install Extensions``` and install Jython. This will require a restart of Ghidra.

Once you run the script it will ask for a file, provide it the output of GoReSym, and now all functions will be properly named.

The start of a Go program will be ```main.main``` and that should exist and should look very similar to your code in the disassembly.

## Getting Go Strings
There are many open source projects to use that will label the strings appropriately, but I advise using the one here. https://github.com/volexity/GoStringExtractor

Just follow the instructions on the README to install it, but keep in mind you'll have to start Ghidra in pyGhidra mode, which can be found under the ```support``` directory.

If you run into problems starting pyGhidra mode, there is likely a problem with the Python virtual environment.

## Go Calling Convention
Go has a specific calling convention that is used across the same processors no matter the OS.

This is the calling convention for x64.

```xml
<prototype name="__golang_regabi" extrapop="8" stackshift="8">
  <input>
    <!-- Float args first (storage class order required by Ghidra) -->
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM0_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM1_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM2_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM3_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM4_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM5_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM6_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM7_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM8_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM9_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM10_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM11_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM12_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM13_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM14_Qa"/>
    </pentry>
    <!-- Integer / pointer args -->
    <pentry minsize="1" maxsize="8">
      <register name="RAX"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="RBX"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="RCX"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="RDI"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="RSI"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="R8"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="R9"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="R10"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="R11"/>
    </pentry>
    <!-- Stack spill -->
    <pentry minsize="1" maxsize="500" align="8">
      <addr offset="8" space="stack"/>
    </pentry>
  </input>

  <output>
    <!-- Float returns -->
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM0_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM1_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM2_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM3_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM4_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM5_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM6_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM7_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM8_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM9_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM10_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM11_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM12_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM13_Qa"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="XMM14_Qa"/>
    </pentry>
    <!-- Integer returns -->
    <pentry minsize="1" maxsize="8">
      <register name="RAX"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="RBX"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="RCX"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="RDI"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="RSI"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="R8"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="R9"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="R10"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="R11"/>
    </pentry>
  </output>

  <unaffected>
    <register name="RBP"/>
    <register name="R12"/>
    <register name="R13"/>
    <register name="R14"/>
    <register name="RSP"/>
  </unaffected>

  <killedbycall>
    <register name="R15"/>
  </killedbycall>

</prototype>
```

This is the calling convention for ARM64.

```xml
<prototype name="__golang_arm64_regabi" extrapop="0" stackshift="0">
  <input>
    <!-- Float args first (storage class ordering required by Ghidra) -->
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f0"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f1"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f2"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f3"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f4"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f5"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f6"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f7"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f8"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f9"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f10"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f11"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f12"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f13"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f14"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f15"/>
    </pentry>
    <!-- Integer / pointer args: R0-R15 -->
    <pentry minsize="1" maxsize="8">
      <register name="x0"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x1"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x2"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x3"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x4"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x5"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x6"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x7"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x8"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x9"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x10"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x11"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x12"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x13"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x14"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x15"/>
    </pentry>
    <!-- Stack spill for remaining arguments -->
    <pentry minsize="1" maxsize="500" align="8">
      <addr offset="0" space="stack"/>
    </pentry>
  </input>

  <output>
    <!-- Float returns: F0-F15 -->
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f0"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f1"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f2"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f3"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f4"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f5"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f6"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f7"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f8"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f9"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f10"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f11"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f12"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f13"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f14"/>
    </pentry>
    <pentry minsize="4" maxsize="8" metatype="float">
      <register name="f15"/>
    </pentry>
    <!-- Integer returns: R0-R15 -->
    <pentry minsize="1" maxsize="8">
      <register name="x0"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x1"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x2"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x3"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x4"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x5"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x6"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x7"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x8"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x9"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x10"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x11"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x12"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x13"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x14"/>
    </pentry>
    <pentry minsize="1" maxsize="8">
      <register name="x15"/>
    </pentry>
  </output>

  <unaffected>
    <!-- Callee-saved: x19-x28, fp (x29), sp -->
    <register name="x19"/>
    <register name="x20"/>
    <register name="x21"/>
    <register name="x22"/>
    <register name="x23"/>
    <register name="x24"/>
    <register name="x25"/>
    <!-- x26 = closure context (REGCTXT), preserved across calls -->
    <register name="x26"/>
    <!-- x27 = linker temp (REGTMP), treated as preserved -->
    <register name="x27"/>
    <!-- x28 = g (goroutine pointer), must survive calls -->
    <register name="x28"/>
    <!-- x29 = frame pointer -->
    <register name="x29"/>
    <register name="sp"/>
    <!-- Callee-saved float regs: f16-f31 lower 64 bits preserved -->
    <register name="f16"/>
    <register name="f17"/>
    <register name="f18"/>
    <register name="f19"/>
    <register name="f20"/>
    <register name="f21"/>
    <register name="f22"/>
    <register name="f23"/>
    <register name="f24"/>
    <register name="f25"/>
    <register name="f26"/>
    <register name="f27"/>
    <register name="f28"/>
    <register name="f29"/>
    <register name="f30"/>
    <register name="f31"/>
  </unaffected>

  <killedbycall>
    <!-- Caller-saved scratch: x16 (IP0), x17 (IP1), x18 (platform reg) -->
    <register name="x16"/>
    <register name="x17"/>
    <register name="x18"/>
    <!-- x30 = link register, clobbered by BL -->
    <register name="x30"/>
  </killedbycall>

</prototype>
```

Write one of these calling conventions to an xml and import it into an open Ghidra instance, with a binary being analyzed. 

``
Edit->Options for 'binary_name'->Specification Extensions->Import.
``

Then apply this calling convention to every function.

Make the following Python script and run it in 'Script Manager'. Make sure to change the name of the calling convention if you are using the ARM one.
```python
from ghidra.app.script import GhidraScript
from ghidra.program.model.listing import Function

def run():
    fm = currentProgram.getFunctionManager()
    functions = fm.getFunctions(True)
    count = 0
    skipped = 0

    for func in functions:
        try:
            func.setCallingConvention("__golang_regabi")
            count += 1
        except Exception as e:
            skipped += 1

    print("Applied Go calling convention to {} functions ({} skipped)".format(count, skipped))

run()
```

## On Your Own
Play around with Ghidra and use it to develop an understanding of how Go is compiled, as well as the mechanisms of Go.

Go does things differently than C and most languages, speculate by reverse engineering in Ghidra. Then you can confirm by research via AI, reading research papers or official Go documentation.

Attempt to build structures for data types in Go, slices and strings. Hint: Go is open source. 

There are also plenty of open source projects to help automate Go reverse engineering, even with types, you can attempt to get those working and show them off to the instructor and peers. If it's good enough maybe it will be added to the next rendition of the course.