"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { signOut } from "@/services/auth";
import { deleteClothing }
from "@/services/deleteClothing";
import { useTransition }
from "react";

const COLORS = [
  { id: 1, name: "Black" },
  { id: 2, name: "White" },
  { id: 3, name: "Gray" },
  { id: 4, name: "Blue" },
  { id: 5, name: "Red" },
  { id: 6, name: "Green" },
  { id: 7, name: "Yellow" },
  { id: 8, name: "Brown" },
  { id: 10, name: "Pink" },
  { id: 11, name: "Purple" },
  { id: 12, name: "Orange" },
];

const STYLES = [
  { id: 1, name: "Casual" },
  { id: 2, name: "Formal" },
  { id: 3, name: "Sporty" },
];

const TYPES = [
  { id: 1, name: "Tshirt" },
  { id: 2, name: "Hoodie" },
  { id: 3, name: "Sweater" },
  { id: 4, name: "Jacket" },
  { id: 5, name: "Coat" },
  { id: 6, name: "Shorts" },
  { id: 7, name: "Skirt" },
  { id: 8, name: "Jeans" },
  { id: 9, name: "Sandals" },
  { id: 10, name: "Sneakers" },
  { id: 11, name: "Boots" },
  { id: 12, name: "Long-sleeve" },
  { id: 13, name: "Top" },
  { id: 14, name: "Overshirt" },
  { id: 15, name: "Shoes" },
  { id: 16, name: "Heels" },
];


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


  const [filterTypeId, setFilterTypeId] =
    useState<number | null>(null);

  const [filterColorId, setFilterColorId] =
    useState<number | null>(null);

  const [filterStyleId, setFilterStyleId] =
    useState<number | null>(null);

  const [selectedItem, setSelectedItem] =
    useState<ClothingItem | null>(
      null
    );
  
    const [searchOpen, setSearchOpen] =
  useState(false);

  const [confirmDelete,
  setConfirmDelete] =
    useState(false);

  const [isPending,
  startTransition] =
    useTransition();

  const filteredItems =
  clothingItems.filter(
    (item) =>
      (!filterTypeId ||
        item.type?.id ===
          filterTypeId) &&
      (!filterColorId ||
        item.color?.id ===
          filterColorId) &&
      (!filterStyleId ||
        item.style?.id ===
          filterStyleId)
  );  
    
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
            onClick={() => {
              
              router.push("/addItem")
            }}
          >
            Add Item
          </button>

            <button
              onClick={() => {
                setSelectedItem(null);
                setSearchOpen(
                  !searchOpen
                );
              }}
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
                {(filteredItems ?? []).map(
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
                      onClick={() =>
                        router.push(
                          `/modifyItem/${selectedItem.id}`
                        )
                      }
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
                        marginTop: "1rem",
                        width: "100%",
                      }}
                      onClick={() => {
                        setFilterTypeId(null);
                        setFilterColorId(null);
                        setFilterStyleId(null);
                      }}
                    >
                      Clear Filters
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