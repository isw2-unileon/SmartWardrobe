"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { generateOutfit } from "@/services/generateOutfit";


export default function CreateOutfitTomorrowForm() {
  const [city, setCity] = useState("");
  const [country, setCountry] = useState("");

  const router = useRouter();

  const handleGenerate = async () => {
    const tomorrow = new Date();

    tomorrow.setDate(tomorrow.getDate() + 1);

    const tomorrowString = tomorrow.toISOString().split("T")[0];

    const result = await generateOutfit({
      city,
      startDate: tomorrowString,
      endDate: tomorrowString,
    });

    localStorage.setItem(
      "generatedOutfit",
      JSON.stringify({
        mode: "tomorrow",
        city,
        startDate: tomorrowString,
        endDate: tomorrowString,
        result,
      }),
    );

    router.push("/createOutfit/result");
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
          <h2>Create Outfit</h2>

          <label>City</label>

          <input
            type="text"
            value={city}
            onChange={(e) => setCity(e.target.value)}
            placeholder="León"
          />

          <>
            <label>Country</label>
            <input type="text" placeholder="Spain" />
          </>

          <button onClick={handleGenerate}>Generate Outfit</button>
        </div>
      </div>
    </div>
  );
}
