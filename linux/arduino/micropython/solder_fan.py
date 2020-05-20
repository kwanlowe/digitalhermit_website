import machine
import time


def toggle(p):
    p.value(not p.value())

disabled=1
enabled=0

pollution_threshold = 600

led1 = machine.Pin(2, machine.Pin.OUT)
led2 = machine.Pin(16, machine.Pin.OUT)
relay = machine.Pin(5, machine.Pin.OUT)
sensor = machine.ADC(0)

led1.value(disabled)
led2.value(disabled)

while True:
    print("Sensor reading:", sensor.read())
    sensor_reading = sensor.read()
    if sensor_reading > pollution_threshold:
        relay.on()
        print("Sensor reading above threshold. Enabling fan.")
        led2.value(enabled)
    else:
        relay.off()
        led2.value(disabled)
        print("Sensor reading below threshold. Disabling fan.")
    time.sleep_ms(8000)
