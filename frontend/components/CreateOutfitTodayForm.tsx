"use client";

import { useState } from "react";

export default function CreateOutfitTodayForm() {
  const [city, setCity] =
    useState("");

  return (
    <div className="page-container">
      <div
        className="card"
        style={{
          width: "100%",
          maxWidth: "500px",
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

        <button>
          Generate Outfit
        </button>
      </div>
    </div>
  );
}