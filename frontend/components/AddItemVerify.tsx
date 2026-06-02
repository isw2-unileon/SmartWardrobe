"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { saveClothingItem } from "@/services/clothing";

const COLORS = [
  { id: 1, name: "Black" },
  { id: 2, name: "White" },
  { id: 3, name: "Gray" },
  { id: 4, name: "Blue" },
  { id: 5, name: "Red" },
  { id: 6, name: "Green" },
  { id: 7, name: "Yellow" },
  { id: 8, name: "Brown" },
  { id: 10, name: "Pink" },
  { id: 11, name: "Purple" },
  { id: 12, name: "Orange" },
];

const STYLES = [
  { id: 1, name: "Casual" },
  { id: 2, name: "Formal" },
  { id: 3, name: "Sporty" },
];

const TYPES = [
  { id: 1, name: "Tshirt" },
  { id: 2, name: "Hoodie" },
  { id: 3, name: "Sweater" },
  { id: 4, name: "Jacket" },
  { id: 5, name: "Coat" },
  { id: 6, name: "Shorts" },
  { id: 7, name: "Skirt" },
  { id: 8, name: "Jeans" },
  { id: 9, name: "Sandals" },
  { id: 10, name: "Sneakers" },
  { id: 11, name: "Boots" },
  { id: 12, name: "Long-sleeve" },
  { id: 13, name: "Top" },
  { id: 14, name: "Overshirt" },
  { id: 15, name: "Shoes" },
  { id: 16, name: "Heels" },
];

export default function AddItemVerify({
  imageUrl,
}: {
  imageUrl: string;
}) {
  const router = useRouter();

  const [colorId, setColorId] = useState(1);
  const [styleId, setStyleId] = useState(1);

  const [typeId, setTypeId] = useState(1);


  const [loading, setLoading] =
    useState(false);

  const handleSave = async () => {
    setLoading(true);

    await saveClothingItem({
      imageUrl,
      colorId,
      styleId,
      typeId,
    });

    router.push("/mainMenu");
  };

  return (
    <div className="page-container">
      <div
        className="card"
        style={{
          width: "100%",
          maxWidth: "950px",
        }}
      >
        <div
          style={{
            display: "grid",
            gridTemplateColumns: "1fr 1fr",
            gap: "2rem",
            alignItems: "start",
          }}
        >
          <div>
            <img
              src={imageUrl}
              alt="preview"
              style={{
                width: "100%",
                height: "420px",
                objectFit: "contain",
                borderRadius: "20px",
                border:
                  "1px solid #E5D8CC",
              }}
            />
          </div>

          <div
            style={{
              display: "flex",
              flexDirection: "column",
              gap: "1.2rem",
            }}
          >
            <h2>Verify Item</h2>
            <label>Type</label>

            <select
            value={typeId}
            onChange={(e) =>
                setTypeId(
                Number(e.target.value)
                )
            }
            >
            {TYPES.map((t) => (
                <option
                key={t.id}
                value={t.id}
                >
                {t.name}
                </option>
            ))}
            </select>

            <label>Color</label>

            <select
              value={colorId}
              onChange={(e) =>
                setColorId(
                  Number(e.target.value)
                )
              }
            >
              {COLORS.map((c) => (
                <option
                  key={c.id}
                  value={c.id}
                >
                  {c.name}
                </option>
              ))}
            </select>

            <label>Style</label>

            <select
              value={styleId}
              onChange={(e) =>
                setStyleId(
                  Number(e.target.value)
                )
              }
            >
              {STYLES.map((s) => (
                <option
                  key={s.id}
                  value={s.id}
                >
                  {s.name}
                </option>
              ))}
            </select>

            <button
              disabled={loading}
              onClick={handleSave}
            >
              {loading
                ? "Saving..."
                : "OK"}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}