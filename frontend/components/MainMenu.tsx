"use client";

import { useRouter } from "next/navigation";

export default function MainMenu() {
  const router = useRouter();

  const mockItems = Array.from({ length: 12 });

  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        padding: "3rem",
      }}
    >
      <div
        style={{
          width: "100%",
          maxWidth: "1100px",
          backgroundColor: "#FFFDFB",
          borderRadius: "28px",
          padding: "2rem",
          boxShadow:
            "0 10px 25px rgba(0,0,0,0.05), 0 4px 10px rgba(0,0,0,0.03)",
          border: "1px solid #E5D8CC",
        }}
      >
        {/* Top section */}

        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            marginBottom: "2rem",
            gap: "1rem",
          }}
        >
          <div
            style={{
              display: "flex",
              gap: "1rem",
            }}
          >
            <button onClick={() => router.push("/add-item")}>
              Add Item
            </button>

            <button onClick={() => router.push("/search-item")}>
              Search Item
            </button>
          </div>

          <button onClick={() => router.push("/login")}>
            Log Out
          </button>
        </div>

        {/* Wardrobe grid */}

        <div
          style={{
            border: "1px solid #E5D8CC",
            borderRadius: "22px",
            padding: "1.5rem",
            height: "420px",
            overflowY: "auto",
            marginBottom: "2rem",
            backgroundColor: "#FCFAF7",
          }}
        >
          <div
            style={{
              display: "grid",
              gridTemplateColumns: "repeat(3, 1fr)",
              gap: "1.5rem",
            }}
          >
            {mockItems.map((_, index) => (
              <div
                key={index}
                style={{
                  aspectRatio: "1",
                  border: "2px solid #C8B6A6",
                  borderRadius: "18px",
                  backgroundColor: "#FFFDFB",
                }}
              />
            ))}
          </div>
        </div>

        {/* Create outfit */}

        <div>
          <h2
            style={{
              marginBottom: "1.5rem",
            }}
          >
            Create Outfit
          </h2>

          <div
            style={{
              display: "flex",
              gap: "1rem",
              flexWrap: "wrap",
            }}
          >
            <button
              onClick={() =>
                router.push("/create-outfit/today")
              }
            >
              Today
            </button>

            <button
              onClick={() =>
                router.push("/create-outfit/tomorrow")
              }
            >
              Tomorrow
            </button>

            <button
              onClick={() =>
                router.push("/create-outfit/week")
              }
            >
              For a Week
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}