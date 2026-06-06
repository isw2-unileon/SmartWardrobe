/* eslint-disable @next/next/no-img-element */
"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { generateOutfit } from "@/services/generateOutfit";

export default function WeekOutfitResult() {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const [data, setData] = useState<any>(null);

  const router = useRouter();
  const handleRetryWeek = async () => {
    if (!data) {
      return;
    }

    const outfits = [];

    const start = new Date(data.startDate);

    for (let i = 0; i < 7; i++) {
      const current = new Date(start);

      current.setDate(start.getDate() + i);

      const date = current.toISOString().split("T")[0];

      const result = await generateOutfit({
        city: data.city,
        country: data.country,
        startDate: date,
        endDate: date,
      });

      outfits.push({
        date,
        result,
      });
    }

    const newData = {
      ...data,
      outfits,
    };

    localStorage.setItem("generatedOutfit", JSON.stringify(newData));

    setData(newData);
  };

  useEffect(() => {
    const stored = localStorage.getItem("generatedOutfit");

    if (!stored) {
      router.push("/mainMenu");
      return;
    }

    if (stored) {
      // eslint-disable-next-line react-hooks/set-state-in-effect
      setData(JSON.parse(stored));
    }
  }, [router]);

  if (!data) {
    return <p>Loading...</p>;
  }

  return (
    <div className="page-container">
      <div
        className="card"
        style={{
          maxWidth: "1400px",
          width: "100%",
          overflowX: "auto",
        }}
      >
        <h2>Weekly Outfits</h2>

        <div
          style={{
            display: "grid",
            gridTemplateColumns: "repeat(7, 1fr)",
            gap: "1rem",
            marginTop: "2rem",
          }}
        >
          {data.outfits.map((day: any, index: number) => {
            const outfit = day?.result?.[0]?.outfit;

            if (!outfit) {
              return null;
            }

            return (
              <div
                key={index}
                style={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  gap: "1rem",
                }}
              >
                <div
                  style={{
                    backgroundColor: "#8B6B4A",
                    color: "white",
                    borderRadius: "12px",
                    padding: "0.5rem 1rem",
                    fontWeight: 600,
                    width: "100%",
                    textAlign: "center",
                  }}
                >
                  {new Date(day.date).toLocaleDateString("es-ES", {
                    day: "2-digit",
                    month: "2-digit",
                  })}
                </div>

                {outfit.upperwear && (
                  <img
                    src={outfit.upperwear.imageUrl}
                    alt=""
                    style={{
                      width: "120px",
                      height: "120px",
                      objectFit: "contain",
                      border: "1px solid #B8A391",
                      borderRadius: "12px",
                    }}
                  />
                )}

                {outfit.bottomwear && (
                  <img
                    src={outfit.bottomwear.imageUrl}
                    alt=""
                    style={{
                      width: "120px",
                      height: "120px",
                      objectFit: "contain",
                      border: "1px solid #B8A391",
                      borderRadius: "12px",
                    }}
                  />
                )}

                {outfit.footwear && (
                  <img
                    src={outfit.footwear.imageUrl}
                    alt=""
                    style={{
                      width: "120px",
                      height: "120px",
                      objectFit: "contain",
                      border: "1px solid #B8A391",
                      borderRadius: "12px",
                    }}
                  />
                )}

                {outfit.outerwear && (
                  <img
                    src={outfit.outerwear.imageUrl}
                    alt=""
                    style={{
                      width: "120px",
                      height: "120px",
                      objectFit: "contain",
                      border: "1px solid #B8A391",
                      borderRadius: "12px",
                    }}
                  />
                )}
              </div>
            );
          })}
        </div>

        <div
          style={{
            display: "flex",
            justifyContent: "center",
            gap: "1rem",
            marginTop: "2rem",
          }}
        >
          <button onClick={handleRetryWeek}>Retry Week</button>

          <button onClick={() => router.push("/mainMenu")}>
            Return to Main Menu
          </button>
        </div>
      </div>
    </div>
  );
}
