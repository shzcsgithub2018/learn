from langchain.llms import OpenAI
from langchain.prompts import PromptTemplate
from langchain.chains import LLMChain

# 设置 OpenAI API 密钥
import os
os.environ["OPENAI_API_KEY"] = "your_openai_api_key"

# 创建一个语言模型实例
llm = OpenAI()

# 创建一个提示模板
prompt_template = "请描述一下 {subject}。"
prompt = PromptTemplate.from_template(prompt_template)

# 创建一个链
chain = LLMChain(llm=llm, prompt=prompt)

# 使用链生成文本
subject = "猫"
description = chain.run(subject)
print(description)