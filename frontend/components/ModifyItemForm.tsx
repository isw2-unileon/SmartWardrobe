"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { updateClothing } from "@/services/updateClothing";

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

type Props = {
  item: {
    id: number;
    imageUrl: string;

    type: {
      id: number;
      name: string;
    };

    color: {
      id: number;
      name: string;
    };

    style: {
      id: number;
      name: string;
    };
  };
};

export default function ModifyItemForm({
  item,
}: Props) {
  const router = useRouter();

  const [typeId, setTypeId] =
    useState(item.type.id);

  const [colorId, setColorId] =
    useState(item.color.id);

  const [styleId, setStyleId] =
    useState(item.style.id);

  const [loading, setLoading] =
    useState(false);

  const typeName =
    TYPES.find(
      (t) => t.id === typeId
    )?.name || "";

  const colorName =
    COLORS.find(
      (c) => c.id === colorId
    )?.name || "";

  const styleName =
    STYLES.find(
      (s) => s.id === styleId
    )?.name || "";

  const handleSave = async () => {
    setLoading(true);

    await updateClothing({
      id: item.id,

      imageUrl:
        item.imageUrl,

      typeId,
      typeName,

      colorId,
      colorName,

      styleId,
      styleName,
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
            gridTemplateColumns:
              "1fr 1fr",
            gap: "2rem",
            alignItems: "start",
          }}
        >
          <div>
            <img
              src={item.imageUrl}
              alt="preview"
              style={{
                width: "100%",
                height: "420px",
                objectFit:
                  "contain",
                borderRadius:
                  "20px",
                border:
                  "1px solid #E5D8CC",
              }}
            />
          </div>

          <div
            style={{
              display: "flex",
              flexDirection:
                "column",
              gap: "1.2rem",
            }}
          >
            <h2>Modify Item</h2>

            <label>Type</label>

            <select
              value={typeId}
              onChange={(e) =>
                setTypeId(
                  Number(
                    e.target.value
                  )
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
                  Number(
                    e.target.value
                  )
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
                  Number(
                    e.target.value
                  )
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
              onClick={
                handleSave
              }
            >
              {loading
                ? "Saving..."
                : "Save Changes"}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}