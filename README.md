
# Zadanie 2 – Budowa i publikacja obrazu Dockera z multi-architekturą oraz skanowaniem CVE

### Autor
Szymon Zięba
 
### Opis projektu
Projekt zawiera aplikację napisaną w języku Go oraz plik Dockerfile, który pozwala na stworzenie obrazu kontenera.  
Celem zadania było przygotowanie pipeline’u w GitHub Actions, który:
- buduje obraz Dockera dla dwóch architektur: `linux/amd64` oraz `linux/arm64`,
- wykorzystuje cache budowania przechowywane w publicznym repozytorium na DockerHub,
- skanuje obraz pod kątem luk bezpieczeństwa (CVE) przy pomocy narzędzia Trivy,
- publikuje obraz tylko w przypadku braku wykrytych luk o poziomie wysokim lub krytycznym,
- publikuje gotowy obraz do publicznego repozytorium GitHub Container Registry (GHCR).
### Struktura repozytorium
- `main.go` – kod źródłowy aplikacji w Go
- `Dockerfile` – dwustopniowy build obsługujący multiarchitekturę
- `.github/workflows/docker-image.yml` – definicja workflow GitHub Actions
- `templates/` – katalog ze szablonami HTML
### Konfiguracja pipeline'u GitHub Actions
Workflow `docker-image.yml` wykonuje kolejno:
1. Checkout repozytorium z kodem źródłowym.
2. Pobranie metadanych do tworzenia tagów (skrót SHA oraz wersja semantyczna na podstawie tagów git).
3. Przygotowanie środowiska QEMU i Buildx do budowania obrazów dla wielu architektur.
4. Logowanie do GHCR oraz DockerHub (w celu korzystania z cache).
5. Budowanie obrazu z wykorzystaniem cache pobieranego i zapisywanego w publicznym repozytorium DockerHub `${{ vars.DOCKERHUB_USERNAME }}/zadanie2-cache`, co znacząco przyspiesza kolejne buildy.
6. Skanowanie obrazu narzędziem Trivy, które wykrywa potencjalne luki o wysokim i krytycznym poziomie.
7. Publikacja obrazu do GHCR tylko w przypadku, gdy skanowanie nie wykryło poważnych zagrożeń
### Tagowanie obrazów i cache
- Obrazy są automatycznie tagowane przez `docker/metadata-action` z wykorzystaniem:
- `sha-<skrót_SHA>` — unikalny identyfikator builda,
- `semver` — na podstawie git tagów w formacie `vX.Y.Z`.
- -   Cache budowania przechowywany jest w publicznym repozytorium DockerHub `${{ vars.DOCKERHUB_USERNAME }}/zadanie2-cache`, co umożliwia szybkie budowanie obrazów na różnych maszynach.

### Wybór narzędzia do skanowania CVE
Zdecydowano się na użycie **Trivy** ze względu na:
- prostą integrację z pipeline CI/CD,
- automatyczne aktualizacje bazy danych CVE,
- możliwość definiowania progów zagrożeń,
- dużą popularność oraz wsparcie w społeczności DevOps.
### Uruchomienie pipeline
- Automatyczne wyzwalanie pipeline’u następuje po wypchnięciu taga git w formacie `vX.Y.Z`, co powoduje zbudowanie i opublikowanie obrazu.
- Workflow można także uruchomić ręcznie z poziomu interfejsu GitHub lub z linii poleceń:

```bash
gh workflow run docker-image.yml
```
