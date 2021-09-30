i = 0

while not(i % 2 != 0 and i <= 100 and i >= 1):
    i = int(input("Masukan bilangan: ")) 
    if i % 2 != 0 and i <= 100 and i >= 1:
        print(f"Selamat! {i} adalah bilangan ganjil antara 1 - 100")