"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

export async function removeBackground(file: File): Promise<File> {
  const cookieStore = await cookies();

  const supabase = createClient(cookieStore);

  const {
    data: { session },
  } = await supabase.auth.getSession();

  const token = session?.access_token;

  const formData = new FormData();

  formData.append("file", file);

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/removeBackground`,
    {
      method: "POST",

      headers: {
        Authorization: `Bearer ${token}`,
      },

      body: formData,
    },
  );

  if (!response.ok) {
    const error = await response.json();

    throw new Error(error.error || "Background removal failed");
  }

  const blob = await response.blob();

  console.log("Blob size:", blob.size);

  console.log("Blob type:", blob.type);

  const processedFile = new File([blob], `nobg-${file.name}.png`, {
    type: "image/png",
  });

  return processedFile;
}
