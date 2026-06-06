/* eslint-disable @next/next/no-img-element */
"use client";

import { useState } from "react";
import { uploadImage } from "@/services/storage";
import { removeBackground } from "@/services/removeBackground";
//import { analyzeClothing } from "@/services/clip";
import { useRouter } from "next/navigation";

export default function AddItemForm() {
  const [file, setFile] = useState<File | null>(null);
  const [preview, setPreview] = useState<string | null>(null);
  const [removingBg, setRemovingBg] = useState(false);
  const [loading, setLoading] = useState(false);
  const router = useRouter();

  const handleFile = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const selected = e.target.files?.[0];

    if (!selected) return;

    setFile(selected);
    preview;

    setRemovingBg(true);

    try {
      const processedFile = await removeBackground(selected);

      setFile(processedFile);

      setPreview(URL.createObjectURL(processedFile));
    } catch (error) {
      console.error(error);

      setPreview(URL.createObjectURL(selected));
    } finally {
      setRemovingBg(false);
    }
  };

  const handleUpload = async () => {
    if (!file) return;

    setLoading(true);

    const formData = new FormData();
    formData.append("file", file);

    const imageUrl = await uploadImage(formData);

    //const prediction = await analyzeClothing(file);

    setLoading(false);

    // router.push(
    //   `/addItem/Verify?` +
    //   `imageUrl=${encodeURIComponent(imageUrl)}` +
    //   `&color=${encodeURIComponent(prediction.color)}` +
    //   `&style=${encodeURIComponent(prediction.style)}` +
    //   `&type=${encodeURIComponent(prediction.type)}`
    // );
    router.push(`/addItem/Verify?imageUrl=${encodeURIComponent(imageUrl)}`);
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
          <p
            style={{
              textAlign: "center",
              fontWeight: 600,
              margin: 0,
              color: "#4B3A2F",
            }}
          >
            Select your clothing item
          </p>

          <input
            id="file-upload"
            type="file"
            accept="image/*"
            onChange={handleFile}
            style={{ display: "none" }}
          />

          <label
            htmlFor="file-upload"
            style={{
              backgroundColor: "#6D8B74",
              color: "white",
              borderRadius: "12px",
              padding: "12px 24px",
              fontWeight: 600,
              cursor: "pointer",
              width: "fit-content",
              alignSelf: "center",
            }}
          >
            Choose Image
          </label>

          <p
            style={{
              textAlign: "center",
              color: "#7B6A5E",
              margin: 0,
              minHeight: "20px",
            }}
          >
            {file ? file.name : "No image selected"}
          </p>

          {removingBg && (
            <div
              style={{
                textAlign: "center",
                padding: "2rem",
              }}
            >
              <div>Removing background...</div>
            </div>
          )}
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
            {loading ? "Uploading..." : "Upload Photo"}
          </button>
        </form>
      </div>
    </div>
  );
}
