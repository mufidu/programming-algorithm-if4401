i = 0

while i % 2 == 0 or i > 100 or i < 1:
    i = int(input("Masukan bilangan: ")) 
else:
    print(f"Selamat! {i} adalah bilangan ganjil antara 1 - 100")