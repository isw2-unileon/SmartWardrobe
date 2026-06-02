"use client";

import { useState } from "react";
import { uploadImage } from "@/services/storage";

export default function AddItemForm() {
  const [file, setFile] = useState<File | null>(null);
  const [preview, setPreview] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleFile = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    const selected = e.target.files?.[0];

    if (!selected) return;

    setFile(selected);
    setPreview(URL.createObjectURL(selected));
  };

  const handleUpload = async () => {
    if (!file) return;

    setLoading(true);

    const formData = new FormData();
    formData.append("file", file);

    await uploadImage(formData);

    setLoading(false);
  };

  return (
    <div className="page-container">
      <div
        className="card"
        style={{
          width: "100%",
          maxWidth: "650px",
        }}
      >
        <h2
          style={{
            textAlign: "center",
            marginBottom: "1.5rem",
          }}
        >
          Add Item
        </h2>

        <form
          style={{
            display: "flex",
            flexDirection: "column",
            gap: "1.5rem",
          }}
        >
          <input
            type="file"
            accept="image/*"
            onChange={handleFile}
          />

          {preview && (
            <div
              style={{
                border: "1px solid #E5D8CC",
                borderRadius: "20px",
                padding: "1rem",
                backgroundColor: "#FCFAF7",
              }}
            >
              <img
                src={preview}
                alt="preview"
                style={{
                  width: "100%",
                  height: "350px",

                  objectFit: "contain",

                  borderRadius: "16px",
                }}
              />
            </div>
          )}

          <button
            type="button"
            disabled={!file || loading}
            onClick={handleUpload}
          >
            {loading
              ? "Uploading..."
              : "Upload Photo"}
          </button>
        </form>
      </div>
    </div>
  );
}