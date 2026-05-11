#!/usr/bin/env bash
set -euo pipefail

out_path=${1:-data/stores.csv}
count=${2:-100}

mkdir -p "$(dirname "$out_path")"

name_prefixes=(
    "Golden" "Urban" "Spicy" "Blue" "Cozy" "Sunny" "Cedar" "Crimson"
    "Amber" "Basil" "River" "Olive" "Harvest" "Juniper" "Luna" "Maple" "Pepper" "Saffron"
  )
name_suffixes=(
    "Bistro" "Kitchen" "Table" "Fork" "Pantry" "Grill"
    "Market" "Diner" "Eatery" "Canteen" "Cafe" "Brasserie" "Tavern" "Kitchenette" "House" "Bar" "Noodle"
  )
cuisines=(
    "italian" "mexican" "thai" "japanese" "indian" "mediterranean" "american"
    "korean" "vietnamese" "chinese" "greek" "turkish" "lebanese" "peruvian" "ethiopian" "caribbean" "french" "spanish"
  )

{
  echo "id,name,primary_cuisine,cuisines,price_tier,rating_avg,rating_count,orders_7d,is_open_now,created_at,latitude,longitude"
  for i in $(seq 1 "$count"); do
    prefix=${name_prefixes[$RANDOM % ${#name_prefixes[@]}]}
    suffix=${name_suffixes[$RANDOM % ${#name_suffixes[@]}]}
    primary=${cuisines[$RANDOM % ${#cuisines[@]}]}

    extra=${cuisines[$RANDOM % ${#cuisines[@]}]}
    if [[ "$extra" == "$primary" ]]; then
      cuisine_list="$primary"
    else
      cuisine_list="$primary|$extra"
    fi

    price_tier=$((1 + RANDOM % 4))
    rating_avg=$(awk -v r="$RANDOM" 'BEGIN { printf "%.2f", 3 + (r % 200) / 100 }')
    rating_count=$((RANDOM % 500))
    orders_7d=$((RANDOM % 800))
    is_open_now=$((RANDOM % 2))

    age_days=$((RANDOM % 120))
    created_at=$(date -u -d "$age_days days ago" +"%Y-%m-%dT%H:%M:%SZ")

    latitude=$(awk -v r="$RANDOM" 'BEGIN { printf "%.6f", 37.70 + r / 32767 * 0.20 }')
    longitude=$(awk -v r="$RANDOM" 'BEGIN { printf "%.6f", -122.52 + r / 32767 * 0.20 }')

    echo ",${prefix} ${suffix} ${i},${primary},${cuisine_list},${price_tier},${rating_avg},${rating_count},${orders_7d},${is_open_now},${created_at},${latitude},${longitude}"
  done
} > "$out_path"

echo "Wrote ${count} stores to ${out_path}"
