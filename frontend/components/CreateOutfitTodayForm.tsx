"use client";

import { useState } from "react";
import { generateOutfit } from "@/services/generateOutfit";
import { useRouter } from "next/navigation";


export default function CreateOutfitTodayForm() {
  const [city, setCity] =
    useState("");

    const router = useRouter();
    
    const handleGenerate = async () => {
    const today =
        new Date()
        .toISOString()
        .split("T")[0];

    const result =
        await generateOutfit({
        city,
        startDate: today,
        endDate: today,
        });

    localStorage.setItem(
        "generatedOutfit",
        JSON.stringify({
        mode: "today",
        city,
        startDate: today,
        endDate: today,
        result,
        })
    );

    router.push(
        "/createOutfit/result"
    );
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
          <h2>
            Create Outfit
          </h2>

          <label>
            City
          </label>

          <input
            type="text"
            value={city}
            onChange={(e) =>
              setCity(
                e.target.value,
              )
            }
            placeholder="Leon"
          />

          {/*
          <label>
            Country
          </label>

          <input
            type="text"
            placeholder="Spain"
          />
          */}

          <button onClick={handleGenerate}>
            Generate Outfit
            </button>
        </div>
      </div>
    </div>
    
    
  );
}