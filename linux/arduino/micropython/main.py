import machine
import time


def toggle(p):
    p.value(not p.value())


pin2 = machine.Pin(2, machine.Pin.OUT)
pin16 = machine.Pin(16, machine.Pin.OUT)

#for index in range(1,16):
#    print(index)
#    pin = machine.Pin(index, machine.Pin.OUT)
#    toggle(pin)
#    time.sleep_ms(500)

toggle(pin16)

while True:
    toggle(pin2)
    toggle(pin16)
    time.sleep_ms(500)
