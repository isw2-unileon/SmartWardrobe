"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export default function OutfitResult() {
  const [data, setData] =
    useState<any>(null);

  const router =
    useRouter();

  useEffect(() => {
    const stored =
      localStorage.getItem(
        "generatedOutfit"
      );

    if (!stored) {
      router.push(
        "/mainMenu"
      );
      return;
    }

    setData(
      JSON.parse(stored)
    );
  }, [router]);

  if (!data) {
    return <p>Loading...</p>;
  }

  const outfit =
    data.result.outfit;

  return (
    <div className="page-container">
      <div
        className="card"
        style={{
          maxWidth: "900px",
          width: "100%",
        }}
      >
        <h2>
          Generated Outfit
        </h2>

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
              <h3>
                Upperwear
              </h3>

              <img
                src={
                  outfit.upperwear
                    .imageUrl
                }
                alt="upperwear"
                width={200}
              />
            </div>
          )}

          {outfit.bottomwear && (
            <div>
              <h3>
                Bottomwear
              </h3>

              <img
                src={
                  outfit.bottomwear
                    .imageUrl
                }
                alt="bottomwear"
                width={200}
              />
            </div>
          )}

          {outfit.footwear && (
            <div>
              <h3>
                Footwear
              </h3>

              <img
                src={
                  outfit.footwear
                    .imageUrl
                }
                alt="footwear"
                width={200}
              />
            </div>
          )}

          {outfit.outerwear && (
            <div>
              <h3>
                Outerwear
              </h3>

              <img
                src={
                  outfit.outerwear
                    .imageUrl
                }
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
          <button
            onClick={() =>
              window.location.reload()
            }
          >
            Retry
          </button>

          <button
            onClick={() =>
              router.push(
                "/mainMenu"
              )
            }
          >
            Return to Main Menu
          </button>
        </div>
      </div>
    </div>
  );
}