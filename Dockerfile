#Etap 1 - Budowanie aplikacji
FROM golang:1.22 AS builder

# Dodanie informacji o autorze zgodnych z OCI
LABEL org.opencontainers.image.authors="Marek Zając"

# Ustawienie katalogu roboczego
WORKDIR /app

# Kopiowanie plików
COPY . .

# Budowa aplikacji
RUN go build -o zadanie2 main.go

#Etap 2 - Uruchomienie aplikacji
FROM debian:bookworm-slim

# Dodanie informacji o autorze zgodnych z OCI
LABEL org.opencontainers.image.authors="Marek Zając"

# Instalacja certyfikatów SSL (dla net/http)
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates curl \
 && rm -rf /var/lib/apt/lists/*

# Ustawienie katalogu roboczego
WORKDIR /app

# Zmienna środowiskowa portu
ENV PORT=5000

# Kopiowanie plików z poprzedniego etapu
COPY --from=builder /app/zadanie2 .
COPY --from=builder /app/templates ./templates

# Sprawdzenie czy aplikacja działa poprawnie
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s CMD curl --fail http://localhost:5000 || exit 1

# Informacja na jakim porcie nasłuchuje informacja
EXPOSE 5000

# Uruchomienie aplikacji
CMD ["./zadanie2"]
