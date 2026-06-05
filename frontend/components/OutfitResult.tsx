"use client";

import { useEffect, useState } from "react";

import { useRouter } from "next/navigation";

import { generateOutfit } from "@/services/generateOutfit";

export default function OutfitResult() {
  const [data, setData] = useState<any>(null);

  const router = useRouter();

  const handleRetry = async () => {
    const stored = localStorage.getItem("generatedOutfit");

    if (!stored) {
      return;
    }

    const parsed = JSON.parse(stored);

    const result = await generateOutfit({
      city: parsed.city,

      startDate: parsed.startDate,

      endDate: parsed.endDate,
    });

    const newData = {
      ...parsed,
      result,
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

    setData(JSON.parse(stored));
  }, [router]);

  if (!data) {
    return <p>Loading...</p>;
  }

  const outfit = data.result.outfit;

  return (
    <div className="page-container">
      <div
        className="card"
        style={{
          maxWidth: "900px",
          width: "100%",
        }}
      >
        <h2>Generated Outfit</h2>

        <div
          style={{
            display: "flex",
            gap: "2rem",
            flexWrap: "wrap",
            marginTop: "2rem",
          }}
        >
          {outfit.upperwear && (
            <div>
              <h3>Upperwear</h3>

              <img
                src={outfit.upperwear.imageUrl}
                alt="upperwear"
                width={200}
              />
            </div>
          )}

          {outfit.bottomwear && (
            <div>
              <h3>Bottomwear</h3>

              <img
                src={outfit.bottomwear.imageUrl}
                alt="bottomwear"
                width={200}
              />
            </div>
          )}

          {outfit.footwear && (
            <div>
              <h3>Footwear</h3>

              <img src={outfit.footwear.imageUrl} alt="footwear" width={200} />
            </div>
          )}

          {outfit.outerwear && (
            <div>
              <h3>Outerwear</h3>

              <img
                src={outfit.outerwear.imageUrl}
                alt="outerwear"
                width={200}
              />
            </div>
          )}
        </div>

        <div
          style={{
            display: "flex",
            gap: "1rem",
            marginTop: "2rem",
          }}
        >
          <button onClick={handleRetry}>Retry</button>

          <button onClick={() => router.push("/mainMenu")}>
            Return to Main Menu
          </button>
        </div>
      </div>
    </div>
  );
}
