import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import i18n from 'i18next'
import {
  Button,
  Image,
  Segmented,
  Upload,
  GetProp,
  UploadFile,
  UploadProps
} from 'antd'
import { PlusOutlined, LogoutOutlined } from '@ant-design/icons'

import { useAuth } from '@/context/authContext'
import SubtitledElement from '@/components/SubtitledElement/SubtitledElement'
import style from './Settings.module.css'

type FileType = Parameters<GetProp<UploadProps, 'beforeUpload'>>[0]

const getBase64 = (file: FileType): Promise<string> =>
  new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = () => resolve(reader.result as string)
    reader.onerror = error => reject(error)
  })

interface UserSettingsProps {
  t: (key: string) => string
}

const UserSettings: React.FC<UserSettingsProps> = ({ t }) => {
  const { user } = useAuth()

  const [previewOpen, setPreviewOpen] = useState(false)
  const [previewImage, setPreviewImage] = useState('')
  const [fileList, setFileList] = useState<UploadFile[]>([
    {
      uid: '-1',
      name: 'image.png',
      status: 'done',
      url: 'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png'
    }
  ])

  const handlePreview = async (file: UploadFile) => {
    if (!file.url && !file.preview) {
      const newFile = {
        ...file,
        preview: await getBase64(file.originFileObj as FileType)
      }
      setFileList(prevList =>
        prevList.map(f => (f.uid === newFile.uid ? newFile : f))
      )
    }

    setPreviewImage(file.url || (file.preview as string))
    setPreviewOpen(true)
  }

  const handleChange: UploadProps['onChange'] = ({ fileList: newFileList }) =>
    setFileList(newFileList)

  const uploadButton = (
    <button style={{ border: 0, background: 'none' }} type="button">
      <PlusOutlined />
      <div style={{ marginTop: 8 }}>Upload</div>
    </button>
  )

  return (
    <div className={style.settingsContainer}>
      <div className={style.userItem}>
        <Upload
          action="https://660d2bd96ddfa2943b33731c.mockapi.io/api/upload"
          listType="picture-circle"
          fileList={fileList}
          onPreview={handlePreview}
          onChange={handleChange}
        >
          {fileList.length >= 1 ? null : uploadButton}
        </Upload>
        {previewImage && (
          <Image
            wrapperStyle={{ display: 'none' }}
            preview={{
              visible: previewOpen,
              onVisibleChange: visible => setPreviewOpen(visible),
              afterOpenChange: visible => !visible && setPreviewImage('')
            }}
            src={previewImage}
          />
        )}
      </div>
      <div className={style.userInformations}>
        <b>{t('pages.settings.userInfos')}</b>
        <SubtitledElement subtitleKey={t('components.input.firstName.label')}>
          {user?.firstname}
        </SubtitledElement>
        <SubtitledElement subtitleKey={t('components.input.lastName.label')}>
          {user?.lastname}
        </SubtitledElement>
        <SubtitledElement subtitleKey={t('components.input.email.label')}>
          {user?.email}
        </SubtitledElement>
      </div>
    </div>
  )
}

const Settings: React.FC = () => {
  const { t } = useTranslation()
  const { logout } = useAuth()

  const switchLanguage = (language: string) => {
    let lang = ''
    switch (language) {
      case 'fr' as string:
        lang = 'fr'
        break
      case 'en' as string:
        lang = 'en'
        break
      default:
        lang = 'fr'
        break
    }
    i18n.changeLanguage(lang)
  }

  return (
    <div className={style.layoutContainer}>
      <UserSettings t={t} />

      <div className={style.settingsContainer}>
        <div className={style.settingsItem}>
          {t('pages.settings.language')}
          <Segmented
            options={[
              { label: t('pages.settings.fr'), value: 'fr' },
              { label: t('pages.settings.en'), value: 'en' }
            ]}
            value={i18n.language}
            onChange={value => switchLanguage(value as string)}
          />
        </div>

        <div className={style.settingsItem}>
          {t('pages.settings.logout')}
          <Button
            type="primary"
            danger
            shape="circle"
            icon={<LogoutOutlined />}
            onClick={() => logout()}
          />
        </div>
      </div>
    </div>
  )
}

export default Settings
