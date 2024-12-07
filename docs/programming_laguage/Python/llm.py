from http import HTTPStatus
import dashscope
import os

def call_with_messages():
    messages = [{"role":"system","content":""},
                {"role":"user","content":"示例： 1 = A  2 = C  3 = E\n提问：找规律， 4 = ？\n"}]

    responses = dashscope.Generation.call(
        model="qwen-max",
        api_key=os.getenv("DASHSCOPE_API_KEY"),
        messages=messages,
        stream=True,
        result_format='message',  # 将返回结果格式设置为 message
        top_p=0.7,
        temperature=0.7,
        enable_search=False
    )

    for response in responses:
        if response.status_code == HTTPStatus.OK:
            print(response)
        else:
            print('Request id: %s, Status code: %s, error code: %s, error message: %s' % (
                response.request_id, response.status_code,
                response.code, response.message
            ))


if __name__ == '__main__':
    call_with_messages()



    