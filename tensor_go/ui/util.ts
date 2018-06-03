import Axios from "axios";

/**
 * @author Calvin Feng
 */


 export async function classifyImageFile(file: File): Promise<string> {
    const formData: FormData = new FormData();
    formData.append("image", file);

    try {
        const result = await Axios.post("api/tf/recognition/", formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });
        return result.data.message;
    } catch(e) {
        throw e;
    }
 }