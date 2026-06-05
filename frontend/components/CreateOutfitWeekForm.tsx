"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { generateOutfit } from "@/services/generateOutfit";

export default function CreateOutfitWeekForm() {
  const [city, setCity] = useState("Leon");

  const [startDate, setStartDate] = useState(
    new Date().toISOString().split("T")[0],
  );

  const [loading, setLoading] = useState(false);

  const router = useRouter();

  const handleGenerate = async () => {
    setLoading(true);

    const outfits = [];

    const start = new Date(startDate);

    for (let i = 0; i < 7; i++) {
      const current = new Date(start);

      current.setDate(start.getDate() + i);

      const date = current.toISOString().split("T")[0];

      const result = await generateOutfit({
        city,
        startDate: date,
        endDate: date,
      });

      outfits.push({
        date,
        result,
      });
    }

    localStorage.setItem(
      "generatedOutfit",
      JSON.stringify({
        mode: "week",
        city,
        startDate,
        outfits,
      }),
    );

    setLoading(false);

    router.push("/createOutfit/weekResult");
  };

  return (
    <div className="page-container">
      <div
        className="card"
        style={{
          width: "100%",
          maxWidth: "500px",
        }}
      >
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            gap: "1rem",
          }}
        >
          <h2>Create Outfit Week</h2>

          <label>City</label>

          <input
            type="text"
            value={city}
            onChange={(e) => setCity(e.target.value)}
            placeholder="León"
          />

          <label>Start Date</label>

          <input
            type="date"
            value={startDate}
            onChange={(e) => setStartDate(e.target.value)}
          />

          <button
            onClick={handleGenerate}
            disabled={!city || !startDate || loading}
          >
            {loading ? "Generating..." : "Generate Week"}
          </button>
        </div>
      </div>
    </div>
  );
}
