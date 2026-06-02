"use client";

import { useRouter } from "next/navigation";
import { signOut } from "@/services/auth";

export default function MainMenu() {
  const router = useRouter();

  const mockItems = Array.from({ length: 24 });

  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        padding: "2rem",
      }}
    >
      {/* LOG OUT OUTSIDE */}

      <div
        style={{
          width: "100%",
          maxWidth: "1150px",
          position: "relative",
        }}
      >
        <div
          style={{
            display: "flex",
            justifyContent: "flex-end",
            marginBottom: "1rem",
          }}
        >
          <form action={signOut}>
            <button type="submit">Log Out</button>
          </form>
        </div>

        {/* MAIN PANEL */}

        <div
          style={{
            backgroundColor: "#C8B6A6",
            borderRadius: "28px",
            padding: "2rem",

            border: "1px solid #B8A391",

            boxShadow:
              "0 10px 25px rgba(0,0,0,0.05), 0 4px 10px rgba(0,0,0,0.03)",

            display: "flex",
            flexDirection: "column",
            gap: "1.8rem",
          }}
        >
          {/* TOP ACTIONS */}

          <div
            style={{
              display: "flex",
              gap: "1rem",
            }}
          >
            <button onClick={() => router.push("/addItem")}>Add Item</button>

            <button onClick={() => router.push("/searchItem")}>
              Search Item
            </button>
          </div>

          {/* WARDROBE */}

          <div
            style={{
              backgroundColor: "#FFFDFB",

              borderRadius: "24px",

              border: "1px solid #B8A391",

              padding: "1.25rem",

              height: "470px",

              overflowY: "auto",
            }}
          >
            <div
              style={{
                display: "grid",

                gridTemplateColumns: "repeat(4, 1fr)",

                gap: "1rem",
              }}
            >
              {mockItems.map((_, index) => (
                <div
                  key={index}
                  style={{
                    aspectRatio: "1",

                    borderRadius: "18px",

                    border: "2px solid #B8A391",

                    backgroundColor: "#FCFAF7",
                  }}
                />
              ))}
            </div>
          </div>

          {/* CREATE OUTFIT */}

          <div>
            <h2
              style={{
                marginBottom: "1rem",
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
              <button onClick={() => router.push("/createOutfit/today")}>
                Today
              </button>

              <button onClick={() => router.push("/createOutfit/tomorrow")}>
                Tomorrow
              </button>

              <button onClick={() => router.push("/createOutfit/week")}>
                For a Week
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
