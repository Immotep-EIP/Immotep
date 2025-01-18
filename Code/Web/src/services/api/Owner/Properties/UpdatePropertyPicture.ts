import callApi from '@/services/api/apiCaller'

const UpdatePropertyPicture = async (id: string, pictureData: string) => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: `owner/properties/${id}/picture`,
      data: JSON.stringify({ data: pictureData }),
    })
  } catch (error) {
    console.error('Error updating property picture:', error)
    throw error
  }
}

export default UpdatePropertyPicture
