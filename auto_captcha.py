try:
    from PIL import Image
except ImportError:
    import Image
from cv2 import cv2 as cv
import pytesseract


imgFile = "captcha.jpg"


def processImg():
    # 先处理噪点
    grayImg = cv.imread(imgFile, cv.IMREAD_GRAYSCALE)
    if (grayImg==None).any():
        print("Load pic err!")
    height = grayImg.shape[0]
    width = grayImg.shape[1]
    # channel = grayImg.shape[2]
    # print(height,width,channel)
    # for k in range(channel):
    # print("\n\n")
    for i in range(height):
        for j in range(width):
            if grayImg[i][j] < 20:
                grayImg[i][j] = 255
    
    # 二值化
    threshold = 190
    for i in range(height):
        for j in range(width):
            if grayImg[i][j] < threshold:
                grayImg[i][j] = 0
            else:
                grayImg[i][j] = 255

    # cv.imshow("test",grayImg)
    # cv.imwrite("s.jpg",grayImg)

    # 处理孤立点
    for i in range(height):
        for j in range(width):
            if grayImg[i][j] == 0:
                if grayImg[i][j+1] == 255 and grayImg[i][j-1] and grayImg[i+1][j] and grayImg[i-1][j] and grayImg[i-1][j-1] and grayImg[i-1][j+1] and grayImg[i+1][j-1] and grayImg[i+1][j+1]:
                    grayImg[i][j] = 255

    return grayImg
    # cv.imwrite(str(k)+"__.jpg", grayImg)


def showImg_cv(img):
    cv.imshow("Img", img)
    cv.waitKey()

def recognize(img):
    strr = pytesseract.image_to_string(img)
    slices = strr.split(" ")
    ret = "".join(slices)
    # ret = ret[:4]
    print(ret)
  

def main():
  img = processImg()
#   showImg_cv(img)
#   showImg_plt()
  recognize(img)
 
if __name__ == "__main__":
    main()