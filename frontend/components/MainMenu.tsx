"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { signOut } from "@/services/auth";
import { deleteClothing }
from "@/services/deleteClothing";
import { useTransition }
from "react";

type ClothingItem = {
  id: number;

  imageUrl: string;

  type: {
    id: number;
    name: string;
  };

  color: {
    id: number;
    name: string;
  };

  style: {
    id: number;
    name: string;
  };
};

export default function MainMenu({
  clothingItems,
}: {
  clothingItems: ClothingItem[];
}) {

  const router = useRouter();

  const [selectedItem, setSelectedItem] =
    useState<ClothingItem | null>(
      null
    );
  const [confirmDelete,
  setConfirmDelete] =
    useState(false);

  const [isPending,
  startTransition] =
    useTransition();
    
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
      <div
        style={{
          width: "100%",
          maxWidth: "1250px",
          position: "relative",
        }}
      >
        {/* LOG OUT */}

        <div
          style={{
            display: "flex",
            justifyContent: "flex-end",
            marginBottom: "1rem",
          }}
        >
                    <form action={signOut}>
          <button type="submit">
            Log Out
          </button>
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
          {/* TOP */}

          <div
            style={{
              display: "flex",
              gap: "1rem",
            }}
          >
          <button
            onClick={() =>
              router.push("/addItem")
            }
          >
            Add Item
          </button>

            <button
              onClick={() =>
                router.push("/searchItem")
              }
            >
              Search Item
            </button>
          </div>

          {/* GRID + PANEL */}

          <div
            style={{
              display: "flex",
              gap: "1.5rem",
              alignItems: "stretch",
            }}
          >
            {/* WARDROBE */}

            <div
              style={{
                flex: 1,
                backgroundColor:
                  "#FFFDFB",
                borderRadius: "24px",
                border:
                  "1px solid #B8A391",
                padding: "1.25rem",
                height: "470px",
                overflowY: "auto",
              }}
            >
              <div
                style={{
                  display: "grid",
                  gridTemplateColumns:
                    "repeat(4,1fr)",
                  gap: "1rem",
                }}
              >
                {(clothingItems??[]).map(
                  (item) => (
                    <div
                      key={item.id}
                      onClick={() =>
                        setSelectedItem(
                          item
                        )
                      }
                      style={{
                        aspectRatio: "1",
                        borderRadius:
                          "18px",
                        border:
                          "2px solid #B8A391",
                        backgroundColor:
                          "#FCFAF7",
                        overflow: "hidden",
                        cursor: "pointer",
                      }}
                    >
                      <img
                        src={
                          item.imageUrl
                        }
                        alt="clothing"
                        style={{
                          width: "100%",
                          height: "100%",
                          objectFit:
                            "cover",
                        }}
                      />
                    </div>
                  )
                )}
              </div>
            </div>

            {/* SIDE PANEL */}

            {selectedItem && (
              <div
                style={{
                  width: "220px",
                  backgroundColor:
                    "#FFFDFB",
                  borderRadius: "24px",
                  border:
                    "1px solid #B8A391",
                  padding: "1.5rem",
                  display: "flex",
                  flexDirection: "column",
                  gap: "1rem",
                }}
              >
                <img
                  src={
                    selectedItem.imageUrl
                  }
                  alt="preview"
                  style={{
                    width: "100%",
                    aspectRatio: "1",
                    objectFit: "cover",
                    borderRadius: "16px",
                  }}
                />

                {!confirmDelete ? (
                  <>
                    <button
                      style={{
                        width: "100%",
                      }}
                    >
                      Modify
                    </button>

                    <button
                      style={{
                        width: "100%",
                      }}
                      onClick={() =>
                        setConfirmDelete(
                          true
                        )
                      }
                    >
                      Remove
                    </button>

                    <button
                      style={{
                        width: "100%",
                      }}
                      onClick={() =>
                        setSelectedItem(
                          null
                        )
                      }
                    >
                      Close
                    </button>
                  </>
                ) : (
                  <>
                    <p
                      style={{
                        textAlign: "center",
                      }}
                    >
                      Are you sure?
                    </p>

                    <button
                      style={{
                        width: "100%",
                      }}
                      disabled={isPending}
                      onClick={() =>
                        startTransition(
                          async () => {
                            await deleteClothing(
                              selectedItem.id
                            );

                            window.location.reload();
                          }
                        )
                      }
                    >
                      Delete
                    </button>

                    <button
                      style={{
                        width: "100%",
                      }}
                      onClick={() =>
                        setConfirmDelete(
                          false
                        )
                      }
                    >
                      Cancel
                    </button>
                  </>
                )}
              </div>
            )}
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
              <button
                onClick={() =>
                  router.push(
                    "/createOutfit/today"
                  )
                }
              >
                Today
              </button>

              <button
                onClick={() =>
                  router.push(
                    "/createOutfit/tomorrow"
                  )
                }
              >
                Tomorrow
              </button>

              <button
                onClick={() =>
                  router.push(
                    "/createOutfit/week"
                  )
                }
              >
                For a Week
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}