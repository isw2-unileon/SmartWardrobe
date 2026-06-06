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

      country: parsed.country,

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
    if (stored) {
      setData(JSON.parse(stored));
    }
  }, [router]);

  if (!data) {
    return <p>Loading...</p>;
  }
  ///TEMP
  if (data.mode === "week") {
    return (
      <div className="page-container">
        <div className="card">
          <h2>Weekly Outfits</h2>

          <p>Generated: {data.outfits.length} outfits</p>
        </div>
      </div>
    );
  }

  console.log("DATA:", data);

  const outfit = data?.result?.[0]?.outfit;

  if (!outfit) {
    return <p>Outfit not found</p>;
  }

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
              <img
                src={outfit.upperwear.imageUrl}
                alt="upperwear"
                width={200}
              />
            </div>
          )}

          {outfit.bottomwear && (
            <div>
              <img
                src={outfit.bottomwear.imageUrl}
                alt="bottomwear"
                width={200}
              />
            </div>
          )}

          {outfit.footwear && (
            <div>
              <img src={outfit.footwear.imageUrl} alt="footwear" width={200} />
            </div>
          )}

          {outfit.outerwear && (
            <div>
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
